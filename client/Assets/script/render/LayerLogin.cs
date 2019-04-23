
using UnityEngine;
using System.Collections;
using UnityEngine.UI;

class LayerLogin : Layer
{
    private ModuleLogin moduleLogin = null;
    private ModuleNetGameServ moduleNetGameServ = null;

    private InputField account;
    public LayerLogin() : base("ui/login")
    {
    }

    protected override bool init()
    {
        account = UT.getComponent<InputField>(rootObject(), "account");
        account.text = SaveData.data.account;

        Button btnLogin = UT.getComponent<Button>(rootObject(), "BtnLogin");
        btnLogin.onClick.AddListener(this.OnBtnClickLogin);

        return true;
    }

    public override void uninit()
    {
        
    }

    public void OnBtnClickLogin()
    {
        string strAccount = account.text;
        SysLog.debug("start login by " + strAccount);

        UT.startCoroutine(StartLogin(strAccount));
    }

    public override void update(float delta)
    {
        

    }

    private IEnumerator StartLogin(string strAccount)
    {
        SaveData.data.account = strAccount;
        WWW requestToken = new WWW("http://127.0.0.1:18080/login/login?account=" + strAccount);
        yield return requestToken;
        SysLog.debug("get token {0}", requestToken.text);
        moduleNetGameServ.startConnect(Config.ins.gameServerHost.ip, Config.ins.gameServerHost.port);
        moduleLogin.requestLogin(requestToken.text);
    }
}
