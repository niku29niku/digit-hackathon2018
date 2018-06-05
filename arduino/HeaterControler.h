// HeaterControler.h

#ifndef _HEATERCONTROLER_h
#define _HEATERCONTROLER_h

#if defined(ARDUINO) && ARDUINO >= 100
	#include "arduino.h"
#else
	#include "WProgram.h"
#endif

#include <PID_v1.h>

//温度制御に関連する定数
#define SENSER_PIN	0
#define	RERAY_IH_ON	11
#define RERAY_IH_OFF 12
#define SENSER_SAMPLE_COUNT	5
#define	PRE_HEAT_TIME	1050		//[ms]
#define	PID_CONTROL_INTERVAL_TIME	500		//[ms]
#define	MINIMUM_HEAT_TIME_RATE	1		//[%]

#define	NIXIE_TX	2
#define	NIXIE_RX	3
#define	NIXIE_BAUD	9600

//ステートマシン関連
enum heater_control_state {
	START,
	HEATER_IS_PRE_HEAT,
	HEATER_IS_ON,
	HEATER_IS_OFF,
	CALCURATE,
	WAIT,
	FINISH
};

//その他変数に関連する変数
extern PID pid_temperature_control;

void heaterControlFunction();
void initializeMemberVariable();
void initializePin();
void initializePID();
void initializeNixieTransmission();
void setCurrentTemperature();
void serialCommandControl();

int split(String data, char delimiter, String *dst);

void functionForCcommand(String optional_command);
void functionForTcommand(String optional_command);
void functionForLcommand(String optional_command);
void functionForEcommand(String optional_command);
void functionForQuestionMarkCommand(String optional_command);

bool getError();
bool getIsHeating();

void controlIH(bool isON);

void debug();

template<typename T>
inline void serialSendWithCheckError(T data)
{
	getError() ? Serial.println("NG") : Serial.println(data);
}

#endif

