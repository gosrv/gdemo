

using System;

public class ModuleChapter : IModule
{
    private ModuleNetGameServ mGameServer = null;

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

    public void requestChallenge()
    {
        netproto.CS_ChapterChallenge msg = new netproto.CS_ChapterChallenge
        {
        };
        this.mGameServer.send(msg);
    }

    public void requestGetPrize()
    {
        netproto.CS_ChapterGetPrize msg = new netproto.CS_ChapterGetPrize
        {
        };
        this.mGameServer.send(msg);
    }

    [MsgRoute]
    public void response(netproto.SC_ChapterChallenge msg)
    {
        SysLog.debug("chapter challenge result " + msg.code);
    }

    [MsgRoute]
    public void response(netproto.SC_ChapterGetPrize msg)
    {
        SysLog.debug("chapter get prize result " + msg.code);
    }
}