
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
        if (Player.playerData.baseInfo.id > 0)
        {
            SceneRoot.instance.switchLayer(new LayerMainpage());
        }
    }

    public class LoginResult{
        public string Code;
        public string Token;
    }

    private IEnumerator StartLogin(string strAccount)
    {
        Player.playerData = new proto.PlayerData();
        Player.playerInfo = new proto.PlayerInfo();

        SaveData.data.account = strAccount;
        WWW requestToken = new WWW("http://127.0.0.1:18080/login/login?account=" + strAccount);
        yield return requestToken;
        SysLog.debug("get token {0}", requestToken.text);
        LoginResult result = JsonUtility.FromJson<LoginResult>(requestToken.text);
        if (result.Code != "ok")
        {
            SysLog.error("login failed {0}", result.Code);
            yield return null;
        }
        moduleNetGameServ.stopConnect();
        moduleNetGameServ.startConnect(Config.ins.gameServerHost.ip, Config.ins.gameServerHost.port);
        moduleLogin.requestLogin(result.Token);
    }
}
