

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

        ParameterInfo[] paramInfos = mMethod.GetParameters();
        mParameterInfo = paramInfos[0];

        if (!typeof(global::ProtoBuf.IExtensible).IsAssignableFrom(mParameterInfo.ParameterType))
        {
            throw new System.Exception("argument is not protobuf");
        }
    }

    public void call(byte[] msg)
    {
        using (System.IO.MemoryStream mem = new System.IO.MemoryStream(msg))
        {
            mMethod.Invoke(mObj, new object[] { ProtoBuf.Serializer.Deserialize(mParameterInfo.ParameterType, mem) });
        }  
    }
}
