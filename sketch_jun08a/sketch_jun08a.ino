const int b = 5;
const int g = 3;
const int r = 6;

void setup() {
  pinMode(r, OUTPUT);
  pinMode(g, OUTPUT);
  pinMode(b, OUTPUT);
}

int r_sec = 60 * 60 * 2 * 1000; // count down time

int lap_time = 60 * 15;
int event = 5;
int i;

int lastEvenet = 10;

bool stop_flg = false;
unsigned long time;

void ease(int delay_ms, int min, int max, void(*func)(int)){
  for(i=min; i<max; i+=1){
    delay(delay_ms);
    func(i);
  }
  for(i=max; i>min; i-=1){
    delay(delay_ms);
    func(i);
  }
}

int el = 0;
void eventLight(){
  el = 0;
  while(el<10){
    ease(10, 10, 200, Blue);
    ease(10, 10, 200, Yellow);
    el++;
  }
}

void timerStop(){
  ease(1, 10, 200, Blue);
  ease(1, 10, 200, Red);
  ease(1, 10, 200, Yellow);
  ease(1, 10, 200, Cyan);
  ease(1, 10, 200, Blue);
  ease(1, 10, 200, Magenta);
}

// around only
void loop() {

  // time start
  time = millis();
  while(time < r_sec){
    time = millis();

    // // 経過時間が15分の倍数ならevent回eventを行う
    // if(((int)(time/1000) % (15 * 60)) == 0){
    //   while(event > 0){
    //     ease(1, 10, 200, Green);
    //     ease(1, 10, 200, Blue);
    //   }
    //   event--;
    // }


    ease(10, 10, 200, Red);

  }

  // time stop last spurt
  while(lastEvenet>0){
    timerStop();
    lastEvenet--;
  }

  // end
  while(true){
    Extinction();
  }

}


void Red(int val){
  analogWrite(r, val);
}

void Green(int val){
  analogWrite(g, val);
}

void Blue(int val){
  analogWrite(b, val);
}

void Yellow(int val){
  analogWrite(r, val);
  analogWrite(g, val);
}

void Magenta(int val){
  analogWrite(r, val);
  analogWrite(b, val);
}

void Cyan(int val){
  analogWrite(g, val);
  analogWrite(b, val);
}

void White(int val){
  analogWrite(r, val);
  analogWrite(g, val);
  analogWrite(b, val);
}

void Extinction(){
  analogWrite(r, 0);
}
