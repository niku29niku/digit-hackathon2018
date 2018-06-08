package com.example.yoshi1125hisa.roastbeefapp;

import android.Manifest;
import android.app.Application;
import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.app.TaskStackBuilder;
import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.content.res.Resources;
import android.graphics.BitmapFactory;
import android.graphics.Color;
import android.net.Uri;
import android.os.Build;
import android.os.CountDownTimer;
import android.provider.Settings;
import android.support.annotation.NonNull;
import android.support.design.widget.Snackbar;
import android.support.v4.app.ActivityCompat;
import android.support.v4.app.NotificationCompat;
import android.support.v4.content.ContextCompat;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.telephony.PhoneNumberFormattingTextWatcher;
import android.text.InputType;
import android.util.Log;
import android.view.KeyEvent;
import android.view.MotionEvent;
import android.view.View;
import android.view.Window;
import android.view.WindowManager;
import android.view.inputmethod.InputMethodManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RelativeLayout;
import android.widget.TextView;
import android.widget.Toast;
import java.text.SimpleDateFormat;
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.time.chrono.Chronology;
import java.util.Locale;
import java.util.stream.IntStream;

import com.firebase.client.Firebase;
import com.google.firebase.auth.FirebaseAuth;
import com.google.firebase.auth.FirebaseUser;
import com.google.firebase.database.ChildEventListener;
import com.google.firebase.database.DataSnapshot;
import com.google.firebase.database.DatabaseError;
import com.google.firebase.database.DatabaseReference;
import com.google.firebase.database.FirebaseDatabase;
import com.google.firebase.database.IgnoreExtraProperties;
import com.google.firebase.database.ValueEventListener;
import com.jakewharton.threetenabp.AndroidThreeTen;
import com.muddzdev.styleabletoastlibrary.StyleableToast;


public class MainActivity extends AppCompatActivity{

    private TextView timerText;
   // private SimpleDateFormat dataFormat = new SimpleDateFormat("mm"+"分"/*:ss.SSS*/, Locale.US);
   private SimpleDateFormat dataFormat = new SimpleDateFormat("HH:mm:ss.SS", Locale.US);
    FirebaseDatabase database = FirebaseDatabase.getInstance();
    DatabaseReference telRefMsg;
    private EditText telNumText;
    private RelativeLayout relativeLayout;
    private InputMethodManager inputMethodManager;
    public static String willEndAt = "2018-06-06T14:29:54+09:00";
    public static Boolean cookStatus; //false
    String channelId = "RoastBeefApp";
    String notificationName = "お肉が完成しました！";


    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        AndroidThreeTen.init(this); // <-

        setContentView(R.layout.activity_main);

        final Context context = getApplicationContext();


        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            NotificationManager manager = (NotificationManager)getSystemService(Context.NOTIFICATION_SERVICE);

            NotificationChannel channel = new NotificationChannel(
                    // 一意のチャンネルID
                    // ここはどこかで定数にしておくのが良さそう
                    channelId,

                    // 設定に表示されるチャンネル名
                    // ここは実際にはリソースを指定するのが良さそう
                    notificationName,

                    // チャンネルの重要度
                    // 重要度によって表示箇所が異なる
                    NotificationManager.IMPORTANCE_DEFAULT
            );

            // 通知時にライトを有効にする
            channel.enableLights(true);
            // 通知時のライトの色
            channel.setLightColor(Color.WHITE);
            // ロック画面での表示レベル
            channel.setLockscreenVisibility(Notification.VISIBILITY_PUBLIC);

            // チャンネルの登録
            manager.createNotificationChannel(channel);
        }

        if (ContextCompat.checkSelfPermission(
                this, Manifest.permission.INTERNET)== PackageManager.PERMISSION_GRANTED){
            // 許可されている時の処理
        }else{
            //許可されていない時の処理
            if (ActivityCompat.shouldShowRequestPermissionRationale(this, Manifest.permission.INTERNET)) {
                //拒否された時 Permissionが必要な理由を表示して再度許可を求めたり、機能を無効にしたりします。
            } else {
                //まだ許可を求める前の時、許可を求めるダイアログを表示します。
                ActivityCompat.requestPermissions(this, new String[]{Manifest.permission.INTERNET}, 0);
            }
        }


//        FirebaseUser user = FirebaseAuth.getInstance().getCurrentUser();//UserInfo
      /*  DatabaseReference ref = FirebaseDatabase.getInstance().getReference("timer");

        ref.addValueEventListener(new ValueEventListener() {
            @Override
            public void onDataChange(DataSnapshot dataSnapshot) {
                //Log.d("ondatachange", dataSnapshot.getRef().push().getKey().toString());
                //  intent.putExtra("postData", dataSnapshot.getValue(Post.class));
                willEndAt = dataSnapshot.getValue(Post.class).toString();
            cookStatus = Boolean.valueOf(String.valueOf(dataSnapshot.getValue(Post.class)));

            }
            @Override
            public void onCancelled(DatabaseError databaseError) {
                StyleableToast.makeText(context, "エラーが発生しました。", Toast.LENGTH_SHORT, R.style.mytoast).show();

            }
        });
  //      ref.push().setValue(post);*/


        long countNumber = 10000;
        // インターバル(更新時間)
        long interval = 1;
        long second = countNumber / 1000;
        long minute = second / 60;
        long hour = minute / 60;


        inputMethodManager = (InputMethodManager) getSystemService(Context.INPUT_METHOD_SERVICE);
        relativeLayout = findViewById(R.id.relativeLayout);
        telRefMsg = database.getReference("tel");
        telNumText = findViewById(R.id.telNumText);

        timerText = findViewById(R.id.timer);
        timerText.setText(dataFormat.format(0));

        //CountDownTimer(long millisInFuture, long countDownInterval);
        final CountDown countDown = new CountDown(countNumber, interval);




     /*if(cookStatus){
           countDown.start();

        } else{
           timerText.setText("未調理");
    }
*/
        //EditTextにリスナーをセット
        telNumText.setOnKeyListener(new View.OnKeyListener() {
            
            //コールバックとしてonKey()メソッドを定義
            @Override
            public boolean onKey(View v, int keyCode, KeyEvent event) {
                //イベントを取得するタイミングには、ボタンが押されてなおかつエンターキーだったときを指定
                if((event.getAction() == KeyEvent.ACTION_DOWN) && (keyCode == KeyEvent.KEYCODE_ENTER)){
                    //キーボードを閉じる
                    inputMethodManager.hideSoftInputFromWindow(telNumText.getWindowToken(), InputMethodManager.RESULT_UNCHANGED_SHOWN);

                    String getTelNum = telNumText.getText().toString();
                    final Post post = new Post(getTelNum);

                    telRefMsg.addValueEventListener(new ValueEventListener() {
                        @Override
                        public void onDataChange(DataSnapshot dataSnapshot) {
                             StyleableToast.makeText(context, "保存しました。", Toast.LENGTH_SHORT, R.style.mytoast).show();
                            telNumText.getEditableText().clear();


                        }
                        @Override
                        public void onCancelled(DatabaseError databaseError) {
                            StyleableToast.makeText(context, "エラーが発生しました。", Toast.LENGTH_SHORT, R.style.mytoast).show();
                        }
                    });


                    telRefMsg.push().setValue(post);
                    return true;
                }

                return false;
            }
        });



       // countDown.start();
/*
 startButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                // 開始
                countDown.start();
            }
        });
        */
    }

    @Override
    protected void onResume() {
        super.onResume();
    }

    // 画面タップ時の処理
    @Override
    public boolean onTouchEvent(MotionEvent event) {

        // キーボードを隠す
        inputMethodManager.hideSoftInputFromWindow(relativeLayout.getWindowToken(), InputMethodManager.HIDE_NOT_ALWAYS);
        // 背景にフォーカスを移す
        relativeLayout.requestFocus();
        return true;

    }

    class CountDown extends CountDownTimer {

        CountDown(long millisInFuture, long countDownInterval) {
            super(millisInFuture, countDownInterval);
        }

        @Override
        public void onFinish() {
            // 完了
            timerText.setText(dataFormat.format(0));

            NotificationManager notificationManager = (NotificationManager) getSystemService(NOTIFICATION_SERVICE);
            Intent notificationIntent = new Intent(MainActivity.this,TelListActivity.class);
            PendingIntent contentIntent = PendingIntent.getActivity(MainActivity.this, 0, notificationIntent, 0);

            Notification notifications = new NotificationCompat.Builder( MainActivity.this)
                    .setSmallIcon(R.drawable.common_google_signin_btn_icon_light) // アイコン
                    .setTicker("Hello") // 通知バーに表示する簡易メッセージ
                    .setWhen(System.currentTimeMillis()) // 時間
                    .setContentTitle("My notification") // 展開メッセージのタイトル
                    .setContentText("Hello Notification!!") // 展開メッセージの詳細メッセージ
                    .setContentIntent(contentIntent) // PendingIntent
                    .build();

            notificationManager.notify(1, notifications);




        }

        // インターバルで呼ばれる
        @Override
        public void onTick(long millisUntilFinished) {
            // 残り時間を分、秒、ミリ秒に分割
            long mm = millisUntilFinished / 1000 / 60;
            long ss = millisUntilFinished / 1000 % 60;
            long ms = millisUntilFinished - ss * 1000 - mm * 1000 * 60;
          //  timerText.setText(String.format("%1$02d"/*:%2$02d.%3$03d*/ ,mm/*, ss, ms*/));
            timerText.setText(String.format("%1$02d:%2$02d.%3$02d" ,mm, ss, ms));
            timerText.setText(dataFormat.format(millisUntilFinished));

        }
    }

    public void look(View v){
        Intent intent = new Intent(this,TelListActivity.class);

        startActivity(intent);
    }
}
