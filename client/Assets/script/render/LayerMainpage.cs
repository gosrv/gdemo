

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerMainpage : Layer
{
    private Text name;
    private Text level;
    private Text exp;
    private Text gold;

    public LayerMainpage() : base("ui/mainpage")
    {

    }

    protected override bool init()
    {
        name = UT.getComponent<Text>(rootObject(), "txtName");
        level = UT.getComponent<Text>(rootObject(), "txtLevel");
        exp = UT.getComponent<Text>(rootObject(), "txtExp");
        gold = UT.getComponent<Text>(rootObject(), "txtGold");
        updateShow();
        return true;
    }

    public override void uninit()
    {

    }


    public override void update(float delta)
    {
        if (Player.playerData.isDirtybaseInfo)
        {
            updateShow();
        }
    } 

    private void updateShow()
    {
        name.text = Player.playerData.baseInfo.name;
        level.text = Player.playerData.baseInfo.level.ToString();
        exp.text = Player.playerData.baseInfo.exp.ToString();
        gold.text = Player.playerData.baseInfo.gold.ToString();
    }
}
