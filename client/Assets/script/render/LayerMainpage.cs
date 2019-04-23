

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerMainpage : Layer
{
    private GameObject[] mTeamMembers = new GameObject[5];

    public LayerMainpage() : base("ui", "mainpage")
    {

    }

    protected override bool init()
    {
        BtnEventListener listener = null;
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "talk");
        listener.setClickCallback(this.OnBtnClickTalk);

        for (int i = 0; i < 5; i++)
        {
            mTeamMembers[i] = UT.getChild(rootObject(), "team/members/mem" + (i + 1));
            mTeamMembers[i].SetActive(true);
        }

        for (int i = 0; i < 11; i++)
        {
            listener = UT.addComponentTo<BtnEventListener>(rootObject(), "funcs/func" + (i + 1));
            listener.setClickCallback(this.OnBtnClickShowPage, i + 1);
        }

        for (int i = 0; i < 5; i++)
        {
            listener = UT.addComponentTo<BtnEventListener>(
                rootObject(), "team/members/mem" + (i + 1));
            listener.setClickCallback(this.OnBtnClickSelectTeamMember, i);
        }

        return true;
    }

    public override void uninit()
    {

    }


    public override void update(float delta)
    {
        
    }

    public void OnBtnClickShowPage(PointerEventData eventData, object data)
    {
        int idx = (System.Int32)data;
        switch (idx)
        {
            case 1:
                LayerGame.ins.switchLayer(new LayerPackHero("ui", "packheros"));
                break;
            case 2:
                addChild(new LayerPackHeroLevelup("ui", "heroslevelup"));
                break;
        }
    }

    public void OnBtnClickSelectTeamMember(PointerEventData eventData, object data)
    {
        int idx = (System.Int32)data;
        LayerSelectTeam layer = new LayerSelectTeam("ui", "selectteam");
        layer.rootObject().SetActive(true);
        layer.setTeamSlot(idx);
        addChild(layer);
    }

    public void OnBtnClickTalk(PointerEventData eventData, object data)
    {
        addChild(new LayerTalk("ui", "chat"));
    }
}
