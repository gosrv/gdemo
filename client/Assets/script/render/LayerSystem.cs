

using System;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerSystem : Layer
{
    private Dictionary<string, GameObject> mSubPlane = new Dictionary<string, GameObject>();

    public LayerSystem() : base("ui", "sysconfig")
    {

    }

    protected override bool init()
    {
        return true;
    }

    public override void uninit()
    {

    }

    public override void update(float delta)
    {

    }

    private void updateBaseInfo()
    {

    }

    public void onToggleFunc(bool value, object data)
    {
        string strv = data as string;

        if (!value)
        {
            mSubPlane[strv].SetActive(false);
        }
        else
        {
            mSubPlane[strv].SetActive(true);
        }        
    }
}
