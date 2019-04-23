

using System.Collections.Generic;
using System.Reflection;

public class MsgRouteMgr
{
    public static MsgRouteMgr instance = new MsgRouteMgr();
    private Dictionary<int, RouteMethod> mRoutes = new Dictionary<int, RouteMethod>();
    public bool addRoute(object obj, MethodInfo method)
    {
        ParameterInfo[] paramInfos = method.GetParameters();
        if (paramInfos == null || paramInfos.Length != 1)
        {
            return false;
        }
        ParameterInfo info = paramInfos[0];
        int id = 0;
        if (mRoutes.ContainsKey(id))
        {
            return false;
        }

        mRoutes.Add(id, new RouteMethod(obj, method));

        return true;
    }

    public RouteMethod getRoute(int id)
    {
        RouteMethod route = null;
        if(!mRoutes.TryGetValue(id, out route))
        {
            return null;
        }

        return route;
    }
}
