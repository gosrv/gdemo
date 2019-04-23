

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerSelectTeam : Layer
{
    private ModulePackHero mModulePackHero = null;
    private GameObject mHeroPrefb;
    private GameObject[] mCardObj = null;
    private int mTeamSlot = -1;

    public LayerSelectTeam(string group, string name) : base(name)
    {
    }

    public void setTeamSlot(int slot)
    {
        mTeamSlot = slot;
    }

    protected override bool init()
    {
        mHeroPrefb = UT.getChild(rootObject(), "hero/lists/Viewport/Content/hero");
        mHeroPrefb.SetActive(false);

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }

    public void OnBtnClickSelect(PointerEventData eventData, object data)
    {
        int idx = (System.Int32)data;
        mParent.removeChild(this);
    }

    public void OnBtnClickClose(PointerEventData eventData, object data)
    {
        mParent.removeChild(this);
    }
}
