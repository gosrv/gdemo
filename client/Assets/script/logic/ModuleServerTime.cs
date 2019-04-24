
public class ModuleServerTime : IModule
{
    private long serverTime = 0;
    private int startTick = 0;

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
        if (Player.playerInfo.isDirtyserverTime)
        {
            serverTime = Player.playerInfo.serverTime;
            startTick = System.Environment.TickCount / 1000;
        }
    }
    /*单位秒*/
    public long getServerTime()
    {
        return serverTime + System.Environment.TickCount / 1000 - startTick;
    }
}
