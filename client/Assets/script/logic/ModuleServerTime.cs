
public class ModuleServerTime : IModule
{
    private int mServerTime = 0;
    private int mStartTick = 0;

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
    /*单位秒*/
    public int getServerTime()
    {
        return mServerTime + System.Environment.TickCount / 1000 - mStartTick;
    }

    public void syncTime(int servertime)
    {
        mServerTime = servertime;
        mStartTick = System.Environment.TickCount / 1000;
    }

    public bool isSameDay(int time)
    {
        int daysecond = 24 * 36000;
        return time / daysecond == getServerTime() / daysecond;
    }
}
