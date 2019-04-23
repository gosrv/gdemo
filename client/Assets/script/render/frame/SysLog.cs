
using UnityEngine;

class SysLog
{
    public static void debug(string log)
    {
        Debug.Log(log);
    }

    public static void error(string log)
    {
        Debug.LogError(log);
    }
}