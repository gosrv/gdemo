

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
    private Text diamond;
    private Layer subLayer;

    public LayerMainpage() : base("ui/mainpage")
    {

    }

    protected override bool init()
    {
        name = UT.getComponent<Text>(rootObject(), "txtName");
        level = UT.getComponent<Text>(rootObject(), "txtLevel");
        exp = UT.getComponent<Text>(rootObject(), "txtExp");
        gold = UT.getComponent<Text>(rootObject(), "txtGold");
        diamond = UT.getComponent<Text>(rootObject(), "txtDiamond");
        updateShow();

        Button btnFunc = UT.getComponent<Button>(rootObject(), "funcs/Viewport/Content/chapter");
        btnFunc.onClick.AddListener(()=> {
            this.switchSubView(new LayerChapter());
        });
        btnFunc = UT.getComponent<Button>(rootObject(), "funcs/Viewport/Content/hero");
        btnFunc.onClick.AddListener(() => {
            this.switchSubView(new LayerHero());
        });
        btnFunc = UT.getComponent<Button>(rootObject(), "funcs/Viewport/Content/equip");
        btnFunc.onClick.AddListener(() => {
            this.switchSubView(new LayerEquip());
        });
        btnFunc = UT.getComponent<Button>(rootObject(), "funcs/Viewport/Content/shop");
        btnFunc.onClick.AddListener(() => {
            this.switchSubView(new LayerShop());
        });
        this.switchSubView(new LayerChapter());
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
        diamond.text = Player.playerData.baseInfo.diamond.ToString();
    }

    private void switchSubView(Layer layer)
    {
        if (subLayer != null)
        {
            this.removeChild(subLayer);
        }
        this.subLayer = layer;
        this.addChild(layer);
    }
}
