
using UnityEngine;

class SysLog
{
    public static void debug(string fmt, params System.Object[] msg)
    {
        Debug.Log(string.Format(fmt, msg));
    }

    public static void error(string fmt, params System.Object[] msg)
    {
        Debug.LogError(string.Format(fmt, msg));
    }
}