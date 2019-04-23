

using System.Reflection;

public class RouteMethod
{
    private object mObj;
    private MethodInfo mMethod;
    private ParameterInfo mParameterInfo;

    public RouteMethod(object obj, MethodInfo method)
    {
        mObj = obj;
        mMethod = method;
    }

    public void call(object msg)
    {
        mMethod.Invoke(mObj, new object[] {msg});
    }
}
