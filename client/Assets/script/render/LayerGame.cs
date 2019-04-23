

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerGame : Layer
{
    private Layer mActiveLayer;
    public static LayerGame ins = null;
    private ModuleBattle mModuleBattle = null;
    private ModuleServerTime mModuleServerTime = null;

    public LayerGame() : base("game")
    {

    }

    protected override bool init()
    {
        BtnEventListener listener = null;
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "bottom/BtnInfo");
        listener.setClickCallback(this.onBtnClickMainpage);
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "bottom/BtnBattleCopy");
        listener.setClickCallback(this.onBtnClickBattleCopy);
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "bottom/BtnFight");
        listener.setClickCallback(this.onBtnClickFight);
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "bottom/BtnShop");
        listener.setClickCallback(this.onBtnClickShop);
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "bottom/BtnFriend");
        listener.setClickCallback(this.onBtnClickFriend);
        listener = UT.addComponentTo<BtnEventListener>(rootObject(), "bottom/BtnHelp");
        listener.setClickCallback(this.onBtnClickHelp);

        onBtnClickMainpage(null, null);

        ins = this;

        return true;
    }

    public override void uninit()
    {
        ins = null;
    }


    public override void update(float delta)
    {

    }

    public void switchLayer(Layer newLayer)
    {
        if (mActiveLayer != null)
        {
            removeChild(mActiveLayer);
        }
        mActiveLayer = newLayer;
        addChild(mActiveLayer);
        mActiveLayer.rootObject().SetActive(true);
    }

    /*系统设置*/
    private void onBtnClickMainpage(PointerEventData eventData, object data)
    {
        switchLayer(new LayerMainpage());
    }

    private void onBtnClickBattleCopy(PointerEventData eventData, object data)
    {
        switchLayer(new LayerBattleCopy());
    }

    private void onBtnClickFight(PointerEventData eventData, object data)
    {
        switchLayer(new LayerFight());
    }

    private void onBtnClickFriend(PointerEventData eventData, object data)
    {
        switchLayer(new LayerFriend());
    }

    private void onBtnClickShop(PointerEventData eventData, object data)
    {
        switchLayer(new LayerShop());
    }

    private void onBtnClickHelp(PointerEventData eventData, object data)
    {
        switchLayer(new LayerHelp());
    }    
}
