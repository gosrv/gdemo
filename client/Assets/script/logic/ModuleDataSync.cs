

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
    public void response(netproto.PlayerData playerData)
    {
        Player.playerData.FromProto(playerData);
        SysLog.debug("merge playerData");
    }

    [MsgRoute]
    public void response(netproto.PlayerInfo playerInfo)
    {
        Player.playerInfo.FromProto(playerInfo);
        SysLog.debug("merge playerInfo");
    }
}
