package com.example.yoshi1125hisa.roastbeefapp;

import com.google.firebase.database.Exclude;

import java.io.Serializable;
import java.util.HashMap;
import java.util.Map;

public class Post implements Serializable{
    private String telNum;

    public Post(String telNum) {
        this.telNum = telNum;

    }

    public String getTelNum() {
        return  telNum;
    }

    public void setTelNum(String telNum){
        this.telNum = telNum;
    }

    }


