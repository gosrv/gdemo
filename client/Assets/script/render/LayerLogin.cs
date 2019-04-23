

using System;
using UnityEngine;
using UnityEngine.EventSystems;
using UnityEngine.UI;

class LayerLogin : Layer
{
    private ModuleLogin mModuleLogin = null;

    private string mLoginId = "";

    public LayerLogin() : base("ui", "login")
    {
    }

    protected override bool init()
    {
        BtnEventListener listener = UT.addComponentTo<BtnEventListener>(rootObject(), "login");
        listener.setClickCallback(this.OnBtnClickLogin);

        mLoginId = PlayerPrefs.GetString("_my_cardgameid");
        if (mLoginId.Length != 0)
        {
            UT.getComponent<InputField>(rootObject(), "name").text = mLoginId;
        }        
                
        return true;
    }

    public override void uninit()
    {
        
    }

    public void OnBtnClickLogin(PointerEventData eventData, object data)
    {
        mLoginId = UT.getComponent<Text>(rootObject(), "name/Text").text;
        SysLog.debug("start login by " + mLoginId);
    }

    public override void update(float delta)
    {

    }
}
