

using System;
using System.Collections.Generic;

public class ModuleTalk : IModule
{
    private ModuleNetGameServ mGameServer = null;
    private const int MAX_TALK_LOG = 15;

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
}