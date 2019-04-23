

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerBattleCopy : Layer
{
    private int mMissionType = -1;
    private GameObject[] mMissionObject = null;
    private GameObject mLevelPrefeb = null;
    private ModuleBattle mModuleBattle = null;

    public LayerBattleCopy() : base("battlecopy")
    {
    }

    protected override bool init()
    {
        mLevelPrefeb = UT.getChild(rootObject(), "missions/Viewport/Content/level");
        mLevelPrefeb.SetActive(false);

        BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "funcs/normal");
        listener.setClickCallback(this.OnBtnClickSelType, 1);

        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "funcs/elite");
        listener.setClickCallback(this.OnBtnClickSelType, 2);

        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "funcs/hero");
        listener.setClickCallback(this.OnBtnClickSelType, 3);

        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "funcs/act");
        listener.setClickCallback(this.OnBtnClickSelType, 4);
        

        OnBtnClickSelType(null, 1);

        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }

    public void OnBtnClickSelType(PointerEventData eventData, object data)
    {
        int type = (System.Int32)data;
        mMissionType = type;

    }

    public void OnBtnClickSelMission(PointerEventData eventData, object data)
    {
        int idx = (System.Int32)data;
    }
}
