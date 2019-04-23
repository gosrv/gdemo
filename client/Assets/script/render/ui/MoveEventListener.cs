using UnityEngine;
using System.Collections;
using UnityEngine.EventSystems;
using System;

public class MoveEventListener : MonoBehaviour, IDragHandler
{
    public delegate void Callback(PointerEventData eventData, object data = null);
    private Callback mCallback = null;
    private object mData = null;

    MoveEventListener()
    {
        //UnityEngine.EventSystems.IDragHandler
    }

    public void setCallback(Callback click, object data = null)
    {
        mCallback = click;
        mData = data;
    }


    // Use this for initialization
    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {

    }

    public void OnDrag(PointerEventData eventData)
    {
        mCallback(eventData, mData);
    }
}