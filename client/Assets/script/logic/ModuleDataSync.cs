

public class ModuleDataSync : IModule
{
    private ModuleServerTime mModuleServerTime = null;

    public override bool init()
    {
        return true;
    }

    public override void uninit()
    {

    }

    public override bool prestart()
    {
        return true;
    }

    public override void update(int delta)
    {

    }

    [MsgRoute]
    public void responsePlayerData(netproto.PlayerData playerData)
    {
        Player.playerData.FromProto(playerData);
    }

    [MsgRoute]
    public void responsePlayerData(netproto.PlayerInfo playerInfo)
    {
        Player.playerInfo.FromProto(playerInfo);
    }
}
