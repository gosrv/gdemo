

using System.Collections.Generic;

public class FrameWork  {
    public static FrameWork instance = new FrameWork();

    private bool installModules()
    {
        List<IModule> modules = new List<IModule>();
        foreach (System.Type ct in typeof(FrameWork).Assembly.GetTypes())
        {
            if (typeof(IModule).IsAssignableFrom(ct) && !ct.IsAbstract)
            {
                modules.Add((IModule)ct.Assembly.CreateInstance(ct.FullName));
            }
        }

        
        foreach(IModule module in modules)
        {
            if (!ModuleMgr.instance.registerModule(module))
            {
                SysLog.debug("module " + module.GetType().Name + " register failed.");
                return false;
            }
        }

        return true;
    }

	public bool init()
    {
        if (!installModules())
        {
            return false;
        }
        if (!ModuleMgr.instance.init())
        {
            SysLog.debug("module mgr init failed.");
            return false;
        }

		if(!SceneRoot.instance.doInit())
		{
			return false;
		}

        return true;
    }

    public void uninit()
    {
        ModuleMgr.instance.unit();
		SceneRoot.instance.uninit();
    }

    public void tick(int delta)
    {        
		SceneRoot.instance.update(delta/1000.0f);
    }

    public void latetick(int delta)
    {
        /*if (Player.data != null)
        {
            Player.data.clearDirty();
        }*/
        ModuleMgr.instance.tick(delta);
    }
}
