

using System;

public class ModuleLogin : IModule
{
    private ModuleNetGameServ mGameServer = null;
    private bool mIsLogined = false;

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