void setup()
{
  Serial.begin(38400); // opens serial port, sets data rate to 9600 bps
}

char incomingByte = 0; // for incoming serial data
String readString = "";

int onGetCommand(String commandLine)
{
  return 0;
}

void loop()
{

  // send data only when you receive data:
  if (Serial.available() > 0)
  {
    // read the incoming byte:
    incomingByte = Serial.read();
    readString += incomingByte;
    if (incomingByte == '\n')
    {
      int result = onGetCommand(readString);
      if (result == 0) {
        Serial.println("OK");
      } else {
        Serial.println("NG");
      }
      readString = "";
    }
  }
}