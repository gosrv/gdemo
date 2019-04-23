using UnityEngine;
using System.Collections;
using UnityEngine.EventSystems;
using System;

public class BtnEventListener : MonoBehaviour, IPointerClickHandler {
    public delegate void OnClick(PointerEventData eventData, object data = null);
    private OnClick mCbOnClick = null;
    private object mData = null;

    BtnEventListener()
    {
        
    }

    public void setClickCallback(OnClick click, object data = null)
    {
        mCbOnClick = click;
        mData = data;
    }

    public void OnPointerClick(PointerEventData eventData)
    {
        mCbOnClick(eventData, mData);
    }

    // Use this for initialization
    void Start () {
	
	}
	
	// Update is called once per frame
	void Update () {
	
	}
}
