

using System.Collections.Generic;
using System.Reflection;

public class MsgRouteMgr
{
    private Dictionary<int, RouteMethod> mRoutes = new Dictionary<int, RouteMethod>();
    public bool addRoute(int id, object obj, MethodInfo method)
    {
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
