package com.example.yoshi1125hisa.roastbeefapp;

import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import java.util.List;

public class PostAdapter extends ArrayAdapter<Post> {

    List<Post> item;

    public PostAdapter(Context context, int resource, List<Post> objects) {
        super(context, resource, objects);
        item = objects;
    }

    @Override
    public int getCount(){
        return  item.size();
    }

    @Override
    public Post getItem(int position){
        return  item.get(position);
    }


        @Override
        public View getView ( int position, View convertView, ViewGroup parent) {
            Post item = getItem(position);
            final ViewHolder viewHolder;

            if (convertView != null) {
                viewHolder = (ViewHolder) convertView.getTag();
            } else {
                convertView = LayoutInflater.from(getContext()).inflate(com.developer.yoshi1125hisa.roastbeefapp.R.layout.listview_item_post, parent, false);

                viewHolder = new ViewHolder();
              viewHolder.telNumText = convertView.findViewById(com.developer.yoshi1125hisa.roastbeefapp.R.id.telNum);
              convertView.setTag(viewHolder);
            }

            viewHolder.telNumText.setText(item.getTelNum());

            return convertView;
        }
    static class ViewHolder{

    TextView telNumText;
   }
}

