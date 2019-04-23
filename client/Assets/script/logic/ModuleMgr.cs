

using System.Collections.Generic;
using System.Reflection;

public class ModuleMgr
{
    public static ModuleMgr instance = new ModuleMgr();

    private Dictionary<string, IModule> mModules = new Dictionary<string, IModule>();

    public bool registerModule(IModule module)
    {
        if (mModules.ContainsKey(module.GetType().Name))
        {
            return false;
        }        

        mModules.Add(module.GetType().Name, module);
        return true;
    }
	
    public T getModule<T>() where T : IModule
    {
        return getModule(typeof(T).Name) as T;
    }

    public IModule getModule(string name)
    {
        IModule module = null;
        if (!mModules.TryGetValue(name, out module))
        {
            return null;
        }

        return module;
    }

    public bool init()
    {
        foreach (IModule module in mModules.Values)
        {
            /*检查元数据*/
            System.Type moduleType = module.GetType();
            /*连接module之间的引用*/
            FieldInfo[] allField = moduleType.GetFields(
                BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Instance | BindingFlags.ExactBinding);        
            foreach(FieldInfo field in allField)
            {
                if (!typeof(IModule).IsAssignableFrom(field.FieldType))
                {
                    continue;
                }

                field.SetValue(module, getModule(field.FieldType.Name));      
            }

            /*检查网络路由消息*/
            MethodInfo[] allMethod = moduleType.GetMethods(
                BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Instance);
            foreach (MethodInfo method in allMethod)
            {
                object[] attrs = method.GetCustomAttributes(typeof(MsgRouteAttribute), false);
                if (attrs == null || attrs.Length == 0)
                {
                    continue;
                }

                if(!MsgRouteMgr.instance.addRoute(module, method))
                {
                    return false;
                }
            }
        }

        foreach (IModule module in mModules.Values)
        {
            if (!module.init())
            {
                return false;
            }
        }

        return true;
    }

    public void unit()
    {
        foreach (IModule module in mModules.Values)
        {
            module.uninit();
        }
        mModules.Clear();
    }

    public void tick(int delta)
    {
        foreach (IModule module in mModules.Values)
        {
            module.update(delta);
        }
    }
}
