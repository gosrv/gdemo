

using System;

public class ModulePackHero : IModule
{
    private ModuleNetGameServ gameServer = null;

    public override bool init()
    {
        return true;
    }

    public override bool prestart()
    {
        return true;
    }

    public override void uninit()
    {
        
    }

    public override void update(int delta)
    {
        
    }

    public void requestDraw(int num)
    {
        netproto.CS_HeroDraw draw = new netproto.CS_HeroDraw
        {
            num = num,
        };
        gameServer.send(draw);
    }

    [MsgRoute]
    public void response(netproto.SC_HeroDraw msg)
    {
        SysLog.debug("hero draw result " + msg.code);
    }
}