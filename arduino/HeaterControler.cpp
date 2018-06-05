#include "HeaterControler.h"
#include <SoftwareSerial.h>

SoftwareSerial Nixie_control = SoftwareSerial(NIXIE_RX, NIXIE_TX);

//通信制御に関連する変数
String last_received_command;

//ステートマシン関連
heater_control_state state;

bool is_error;
bool is_command_received;

double temperature_function_A;
double temperature_function_B;

double current_temperature;
double target_temperature;

uint16_t start_time;
uint16_t target_time;

unsigned long change_heater_control_time;

bool is_heating;
bool is_target_temperature_changed;
bool is_target_time_changed;

//PID関連に関連する変数
double pid_input, pid_output;
int pid_control_interval_time;

PID pid_temperature_control(&pid_input, &pid_output, &target_temperature, 320, 32, 15, DIRECT);

void debugGetmillis() {
	Serial.print(millis());
	Serial.print("\t");
	Serial.print(state);
	Serial.print("\t");
	Serial.println(change_heater_control_time);
}

void heaterControlFunction() {
	//debugGetmillis();
	switch (state)
	{
	case START:
		state = CALCURATE;
		start_time = millis();
		break;

	case HEATER_IS_PRE_HEAT:
		if (millis() >= change_heater_control_time) {
			state = HEATER_IS_ON;
		}
		else {
			state = HEATER_IS_PRE_HEAT;
			controlIH(true);
		}
		break;

	case HEATER_IS_ON:
		state = WAIT;
		change_heater_control_time = millis() + pid_output;
		controlIH(true);
		break;

	case HEATER_IS_OFF:
		state = WAIT;
		change_heater_control_time = millis() + (PID_CONTROL_INTERVAL_TIME - pid_output);
		controlIH(false);
		break;

	case CALCURATE:
		setCurrentTemperature();
		pid_temperature_control.Compute();
		if (pid_output > PID_CONTROL_INTERVAL_TIME / 100 * MINIMUM_HEAT_TIME_RATE) {
			state = HEATER_IS_PRE_HEAT;
			change_heater_control_time = millis() + PRE_HEAT_TIME;
		}
		else {
			state = WAIT;
			change_heater_control_time = millis() + PID_CONTROL_INTERVAL_TIME;
		}
		break;

	case WAIT:
		if (millis() - start_time > target_time) {
			state = FINISH;
			break;
		}
		else if (millis() >= change_heater_control_time) {
			if (getIsHeating()) {
				state = HEATER_IS_OFF;
			}
			else {
				state = CALCURATE;
			}
		}
		else {
			state = WAIT;
		}
		break;

	case FINISH:
		state = START;
		is_heating = false;
		controlIH(false);
		break;

	default:
		break;
	}
}

void controlNixie(int number) {
	char buff[10] = "";
	sprintf(buff, "%03d", number);
	Nixie_control.println(buff);
}

void initializeMemberVariable()
{
	last_received_command = "";

	is_error = false;
	is_command_received = false;

	temperature_function_A = -21.7777311;
	temperature_function_B = 0.1357988529;

	current_temperature = 0.0;
	target_temperature = 60;

	start_time = 0;
	target_time = 0;
	
	is_heating = false;
	is_target_temperature_changed= false;
	is_target_time_changed= false;

	pid_control_interval_time = 500;
	state = START;
}
void initializePin() {
	pinMode(RERAY_IH_ON, OUTPUT);
	pinMode(RERAY_IH_OFF, OUTPUT);
}
void initializePID()
{
	pid_temperature_control.SetOutputLimits(0, pid_control_interval_time);
	pid_temperature_control.SetMode(AUTOMATIC);
}

void initializeNixieTransmission()
{
	Nixie_control.begin(NIXIE_BAUD);
	controlNixie(0);
}

void setCurrentTemperature()
{
	double input_analog_value = 0;

	for (int i = 0; i < SENSER_SAMPLE_COUNT; i++) {
		input_analog_value += analogRead(SENSER_PIN);
	}
	input_analog_value /= SENSER_SAMPLE_COUNT;

	current_temperature 
		= (input_analog_value * temperature_function_B) + temperature_function_A;
	pid_input = current_temperature;
}

void serialCommandControl()
{
	last_received_command = Serial.readStringUntil('\r');
	Serial.readStringUntil('\n');		//frush '\n'
	String splited_command[4] = { "\0" };
	split(last_received_command, ':', splited_command);

	if (getIsHeating()) {
		switch (splited_command[0].charAt(0))
		{
		case 'E':
			is_heating = false;
			serialSendWithCheckError("OK");
			break;

		case '?':
			functionForQuestionMarkCommand(splited_command[1]);
			break;

		default:
			serialSendWithCheckError("NG");
			break;
		}
	}
	else {
		switch (splited_command[0].charAt(0))
		{
		case 'C':
			functionForCcommand(splited_command[1]);
			break;

		case 'T':
			functionForTcommand(splited_command[1]);
			break;

		case 'L':
			functionForLcommand(splited_command[1]);
			break;

		case 'E':
			functionForEcommand(splited_command[1]);
			break;

		case '?':
			functionForQuestionMarkCommand(splited_command[1]);
			break;

		default:
			serialSendWithCheckError("NG");
			break;
		}
	}
}

int split(String data, char delimiter, String *dst)
{
	int index = 0;
	int arraySize = sizeof(data) / sizeof(data[0]);
	int datalength = data.length();
	for (int i = 0; i < datalength; i++) {
		char tmp = data.charAt(i);
		if (tmp == delimiter) {
			index++;
			if (index >(arraySize - 1)) {
				return -1;
			}
		}
		else dst[index] += tmp;
	}
	return(index + 1);
}

void functionForCcommand(String optional_command) {
	double received_temperature;
	received_temperature = optional_command.toDouble();
	target_temperature = received_temperature;
	serialSendWithCheckError(target_temperature);
}

void functionForTcommand(String optional_command)
{
	int received_time;
	received_time = optional_command.toInt();
	target_time = received_time * 1000;
	serialSendWithCheckError(received_time);
}

void functionForLcommand(String optional_command)
{
	is_heating = true;
	serialSendWithCheckError("OK");
}

void functionForEcommand(String optional_command)
{
	is_heating = false;
	controlIH(false);
	serialSendWithCheckError("OK");
}

void functionForQuestionMarkCommand(String optional_command) {
	if (optional_command == "\0") {
		serialSendWithCheckError("OK");
	}
	else if (optional_command == "C") {
		serialSendWithCheckError(target_temperature);
	}
	else if (optional_command == "T") {
		serialSendWithCheckError(target_time / 1000);
	}
	else if (optional_command == "CC") {
		serialSendWithCheckError(current_temperature);
	}
	else if (optional_command == "CT") {
		serialSendWithCheckError((millis() - start_time) / 1000);
	}
	else if (optional_command == "L") {
		if (getIsHeating()) {
			serialSendWithCheckError("NG");
		}
		else {
			serialSendWithCheckError("OK");
		}
	}
	else {
		serialSendWithCheckError("NG");
	}
}

bool getError()
{
	return is_error;
}
bool getIsHeating() {
	return is_heating;
}

void controlIH(bool isON)
{
	if (isON) {
		digitalWrite(RERAY_IH_OFF, LOW);
		delay(1);
		digitalWrite(RERAY_IH_ON, HIGH);
	}
	else {
		digitalWrite(RERAY_IH_ON, LOW);
		delay(1);
		digitalWrite(RERAY_IH_OFF, HIGH);
	}
}
void debug() {
	Serial.println("temp\tpower");
	Serial.print(current_temperature);
	Serial.print(",");
	Serial.print(pid_output / pid_control_interval_time * 100);
	Serial.println("");
}