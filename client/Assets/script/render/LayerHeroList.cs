

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerHeroList : Layer
{
    private GameObject mHeroPrefb;
    private GameObject[] mCardObj = null;

    public LayerHeroList()
    {
    }

    protected override bool init()
    {
        mHeroPrefb = UT.getChild(rootObject(), "lists/Viewport/Content/hero");
        mHeroPrefb.SetActive(false);

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {
        
    }
}
