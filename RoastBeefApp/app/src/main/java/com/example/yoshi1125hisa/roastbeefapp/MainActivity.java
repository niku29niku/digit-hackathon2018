package com.example.yoshi1125hisa.roastbeefapp;

import android.annotation.SuppressLint;
import android.app.Notification;
import android.app.NotificationChannel;
import android.app.NotificationManager;
import android.app.PendingIntent;
import android.content.Context;
import android.content.Intent;
import android.graphics.Color;
import android.os.Build;
import android.os.Bundle;
import android.os.CountDownTimer;
import android.support.annotation.NonNull;
import android.support.annotation.Nullable;
import android.support.v4.app.NotificationCompat;
import android.support.v4.app.NotificationManagerCompat;
import android.support.v4.app.TaskStackBuilder;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.TextView;

import com.google.firebase.database.DataSnapshot;
import com.google.firebase.database.DatabaseError;
import com.google.firebase.database.DatabaseReference;
import com.google.firebase.database.FirebaseDatabase;
import com.google.firebase.database.IgnoreExtraProperties;
import com.google.firebase.database.ValueEventListener;

import java.text.DateFormat;
import java.text.ParseException;
import java.util.Date;
import java.util.concurrent.TimeUnit;

import info.vividcode.time.iso8601.Iso8601ExtendedOffsetDateTimeFormat;

public class MainActivity extends AppCompatActivity{

    @IgnoreExtraProperties
    public static class FirebaseTimer {
        public String willEndAt;
        public boolean cooking;

        public FirebaseTimer() {
        }

        public FirebaseTimer(String willEndAt, boolean cooking) {
            this.willEndAt = willEndAt;
            this.cooking = cooking;
        }
    }

    private static final String channelId = "RoastBeefApp";
    TextView timerView;
    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        timerView = findViewById(R.id.timer);
        final NotificationManager manager = (NotificationManager) getSystemService(Context.NOTIFICATION_SERVICE);
        if (manager == null) {
            throw new RuntimeException("NotificationManager is null");
        }

        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.O) {
            final NotificationChannel channel = new NotificationChannel(channelId, "肉", NotificationManager.IMPORTANCE_HIGH);
            channel.setLockscreenVisibility(Notification.VISIBILITY_PUBLIC);
            manager.createNotificationChannel(channel);
        }
        final DatabaseReference timerDatabase = FirebaseDatabase.getInstance().getReference("timer");
        timerDatabase.addValueEventListener(new ValueEventListener() {
            @Override
            public void onDataChange(@NonNull DataSnapshot dataSnapshot) {
                final FirebaseTimer timer = dataSnapshot.getValue(FirebaseTimer.class);
                if (timer != null) {
                    if (timer.cooking) {
                        onCookingStarted(timer);
                    } else {
                        onCookingNotStarted(timer);
                    }
                }
            }

            @Override
            public void onCancelled(@NonNull DatabaseError databaseError) {
            }
        });

        findViewById(R.id.lookButton).setOnClickListener(new View.OnClickListener() {

            @Override
            public void onClick(View view) {
                final Intent intent = new Intent(MainActivity.this, TelListActivity.class);
                startActivity(intent);
            }
        });

    }

    private void onCookingNotStarted(@SuppressWarnings("unused") FirebaseTimer timer) {
        timerView.setText(R.string.waiting_cooking);
    }

    private void onCookingStarted(FirebaseTimer timer) {
        final DateFormat format = new Iso8601ExtendedOffsetDateTimeFormat();
        try {
            final Date willEndAt = format.parse(timer.willEndAt);
            final long duration = willEndAt.getTime() - System.currentTimeMillis();
            new Countdown(duration, 10).start();
        } catch (ParseException e) {
            e.printStackTrace();
        }
    }

    class Countdown extends CountDownTimer {

        Countdown(long millisInFuture, long countDownInterval) {
            super(millisInFuture, countDownInterval);
        }

        @Override
        public void onTick(long l) {
            final long secondsDuration = TimeUnit.MILLISECONDS.toSeconds(l);
            final long minutes = TimeUnit.SECONDS.toMinutes(secondsDuration);
            final long seconds = secondsDuration % 60;
            @SuppressLint("DefaultLocale") final String text = String.format("%02d:%02d", minutes, seconds);
            timerView.setText(text);
        }

        @Override
        public void onFinish() {
            final NotificationManager manager = (NotificationManager)getSystemService(Context.NOTIFICATION_SERVICE);
            if (manager == null) {
                return;
            }
            final Notification notification = new NotificationCompat.Builder(MainActivity.this, channelId)
                    .setContentTitle("ローストビーフ")
                    .setContentText("お肉ができました！！")
                    .setSmallIcon(R.mipmap.ic_launcher_round)
                    .build();
            manager.notify(1, notification);
        }
    }

}
