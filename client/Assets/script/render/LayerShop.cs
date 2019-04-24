

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerShop : Layer
{
    private ModuleShop moduleShop = null;

    public LayerShop() : base("ui/shop")
    {                
    }

    protected override bool init()
    {
        Text txt = UT.getComponent<Text>(rootObject(), "btnGetDiamond1/Text");
        txt.text = "钻石x" + TableMgr.ins.tableCommon[1].Val;
        txt = UT.getComponent<Text>(rootObject(), "btnGetDiamond2/Text");
        txt.text = "钻石x" + TableMgr.ins.tableCommon[2].Val;
        txt = UT.getComponent<Text>(rootObject(), "btnGetDiamond3/Text");
        txt.text = "钻石x" + TableMgr.ins.tableCommon[3].Val;
        txt = UT.getComponent<Text>(rootObject(), "btnGetDiamond4/Text");
        txt.text = "钻石x" + TableMgr.ins.tableCommon[4].Val;
        txt = UT.getComponent<Text>(rootObject(), "btnGetGold/Text");
        txt.text = "金币x" + (TableMgr.ins.tableCommon[5].Val* TableMgr.ins.tableCommon[6].Val);

        Button btn = UT.getComponent<Button>(rootObject(), "btnGetDiamond1");
        btn.onClick.AddListener(() =>
        {
            moduleShop.requestBuyDiamond(0);
        });
        btn = UT.getComponent<Button>(rootObject(), "btnGetDiamond2");
        btn.onClick.AddListener(() =>
        {
            moduleShop.requestBuyDiamond(1);
        });
        btn = UT.getComponent<Button>(rootObject(), "btnGetDiamond3");
        btn.onClick.AddListener(() =>
        {
            moduleShop.requestBuyDiamond(2);
        });
        btn = UT.getComponent<Button>(rootObject(), "btnGetDiamond4");
        btn.onClick.AddListener(() =>
        {
            moduleShop.requestBuyDiamond(3);
        });
        btn = UT.getComponent<Button>(rootObject(), "btnGetGold");
        btn.onClick.AddListener(() =>
        {
            moduleShop.requestBuyGold();
        });
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(float delta)
    {

    }    
}
