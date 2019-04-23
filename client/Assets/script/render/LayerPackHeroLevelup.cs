

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerPackHeroLevelup : Layer
{
    private ModulePackHero mModulePackHero = null;
    private GameObject mHeroPrefb;
    private GameObject[] mCardObj = null;

    public LayerPackHeroLevelup(string group, string name) : base(name)
    {
    }

    protected override bool init()
    {
        mHeroPrefb = UT.getChild(rootObject(), "lists/Viewport/Content/hero");
        mHeroPrefb.SetActive(false);
       
        BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "close");
        listener.setClickCallback(this.OnBtnClickClose);

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }
    
    public void OnBtnClickLevelup(PointerEventData eventData, object data)
    {
        int idx = (System.Int32)data;
    }

    public void OnBtnClickClose(PointerEventData eventData, object data)
    {
        mParent.removeChild(this);
    }
}
