

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerShop : Layer
{
    private ModuleShop mModuleShop = null;

    public LayerShop() : base("shop")
    {
    }

    protected override bool init()
    {
        for (int i = 0; i < 2; i++)
        {
            BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "hero/item" + (i + 1) + "/buy");
            listener.setClickCallback(this.onBtnHero, i + 1);
        }

        for (int i = 0; i < 3; i++)
        {
            BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "shop/item" + (i + 1) + "/buy");
            listener.setClickCallback(this.onBtnShop, i + 1);
        }

        for (int i=0; i<6; i++)
        {
            BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "charge/item"+(i+1)+"/buy");
            listener.setClickCallback(this.onBtnCharge, i);
        }
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }

    public void onBtnHero(PointerEventData eventData, object data)
    {
        
    }

    public void onBtnShop(PointerEventData eventData, object data)
    {
        
    }

    public void onBtnCharge(PointerEventData eventData, object data)
    {        

    }
}
