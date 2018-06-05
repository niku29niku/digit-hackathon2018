#include "HeaterControler.h"

#define	BAUDRATE	38400

void setup()
{
	Serial.begin(BAUDRATE);

	initializePin();
	initializeMemberVariable();
	initializePID();
	initializeNixieTransmission();
	setCurrentTemperature();
}

void loop() {
	if (Serial.available()) {
		serialCommandControl();
	}

	if (getIsHeating()) {
		heaterControlFunction();
		debug();
	}
}