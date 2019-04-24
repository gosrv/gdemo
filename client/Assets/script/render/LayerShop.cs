

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
