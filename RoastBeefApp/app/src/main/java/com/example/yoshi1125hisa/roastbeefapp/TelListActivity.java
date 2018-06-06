package com.example.yoshi1125hisa.roastbeefapp;

import android.content.Context;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.AdapterView;
import android.widget.ArrayAdapter;
import android.widget.ListView;
import android.widget.TextView;
import android.widget.Toast;
import android.widget.Toolbar;

import com.google.firebase.auth.FirebaseAuth;
import com.google.firebase.database.ChildEventListener;
import com.google.firebase.database.DataSnapshot;
import com.google.firebase.database.DatabaseError;
import com.google.firebase.database.DatabaseReference;
import com.google.firebase.database.FirebaseDatabase;
import com.google.firebase.database.ValueEventListener;
import com.muddzdev.styleabletoastlibrary.StyleableToast;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class TelListActivity extends AppCompatActivity {

    ListView listView;
    public static String key = "";
    public static String telNum = "";
    public static String telNumList[] = {""};


   // private String[] telNum = {""};
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_tel_list);

        final Context context = getApplicationContext();
            // Add items via the Button and EditText at the bottom of the view.

// ListViewにArrayAdapterを設定する
        ListView listView = findViewById(R.id.listView);

        // ListViewに表示するリスト項目をArrayListで準備する
        ArrayList arrayList = new ArrayList<>();
        // 要素の削除、順番変更のためArrayListを定義


        arrayList.add("データが登録されていません");

        // リスト項目とListViewを対応付けるArrayAdapterを用意する
        final ArrayAdapter telArrayAdapter = new ArrayAdapter<>(this, android.R.layout.simple_list_item_1, arrayList);
        listView.setAdapter(telArrayAdapter);


        // リスト項目を長押しクリックした時の処理
        listView.setOnItemLongClickListener(new AdapterView.OnItemLongClickListener() {
            /**
             * @param parent   ListView
             * @param view     選択した項目
             * @param position 選択した項目の添え字
             * @param id       選択した項目のID
             */

            public boolean onItemLongClick(AdapterView<?> parent, View view, int position, long id) {
                String deleteItem = (String) ((TextView) view).getText();


                // 選択した項目を削除する
               telArrayAdapter.remove(deleteItem);

                StyleableToast.makeText(context, "登録されていた電話番号を削除しました。", Toast.LENGTH_SHORT, R.style.mytoast).show();
                return false;
            }


        });


        // 受信側で取得したKeyを検索にかけてXYのInchを取得。
        DatabaseReference sendsRef = FirebaseDatabase.getInstance().getReference("tel");
        sendsRef.addListenerForSingleValueEvent(new ValueEventListener() {
            @Override
            public void onDataChange(DataSnapshot snapshot) {
                for (DataSnapshot dSnapshot : snapshot.getChildren()) {
                    // getApplication()でアプリケーションクラスのインスタンスを拾う
                    key = dSnapshot.getKey();
                    telNum = (String) dSnapshot.child("telNum").getValue();



                }
                // 保存した情報を用いた描画処理などを記載する。

            }


            @Override
            public void onCancelled(DatabaseError databaseError) {

            }

        });

        sendsRef.addChildEventListener(new ChildEventListener() {
            @Override
            public void onChildAdded(DataSnapshot dataSnapshot, String s) {
                key = dataSnapshot.getKey();
                telNum = (String) dataSnapshot.child("telNum").getValue();

                // 追加されたTodoのkey、title、isDoneが取得できているので、
                // 保持しているデータの更新や描画処理を行う。
            }
            @Override
            public void onChildChanged(DataSnapshot dataSnapshot, String s) {
                // Changed
            }
            @Override
            public void onChildRemoved(DataSnapshot dataSnapshot) {
                // Removed
            }
            @Override
            public void onChildMoved(DataSnapshot dataSnapshot, String s) {
                // Moved
            }
            @Override
            public void onCancelled(DatabaseError databaseError) {
                // Error
            }
        });
    }
}
