

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerHeroSell : Layer
{
    private ModulePackHero mModulePackHero = null;
    private GameObject mHeroPrefb;    
    private GameObject[] mCardObj = null;

    public LayerHeroSell()
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

    public void OnBtnClickSell(PointerEventData eventData, object data)
    {
        int idx = (System.Int32)data;
    }
}
