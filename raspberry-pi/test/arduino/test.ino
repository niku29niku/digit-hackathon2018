void setup()
{
  Serial.begin(38400); // opens serial port, sets data rate to 9600 bps
}

char incomingByte = 0; // for incoming serial data
String readString = "";

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
      Serial.println("OK");
      readString = "";
    }
  }
}