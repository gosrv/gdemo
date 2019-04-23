using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

public class ProtoUtil
{
    public static int DefVal(int i)
    {
        return 0;
    }

    public static double DefVal(double v)
    {
        return 0;
    }

    public static float DefVal(float v)
    {
        return 0;
    }

    public static long DefVal(long v)
    {
        return 0;
    }

    public static ulong DefVal(ulong v)
    {
        return 0;
    }
    public static uint DefVal(uint v)
    {
        return 0;
    }
    public static bool DefVal(bool v)
    {
        return false;
    }
    public static string DefVal(string v)
    {
        return "";
    }
    public static byte[] DefVal(byte[] v)
    {
        return null;
    }

    public static int OptVal(int? v)
    {
        return v ?? 0;
    }

    public static double OptVal(double? v)
    {
        return v ?? 0;
    }

    public static float OptVal(float? v)
    {
        return v ?? 0;
    }

    public static long OptVal(long? v)
    {
        return v ?? 0;
    }

    public static ulong OptVal(ulong? v)
    {
        return v ?? 0;
    }
    public static uint OptVal(uint? v)
    {
        return v ?? 0;
    }
    public static bool OptVal(bool? v)
    {
        return v ?? false;
    }
    public static string OptVal(string v)
    {
        return v;
    }

    public static void trimList<T>(List<T> list)
    {
        while(list.Count > 0)
        {
            if (list[list.Count-1] != null)
            {
                break;
            }
            list.RemoveAt(list.Count - 1);
        }
    }

    public interface DirtyListener<T>
    {
        bool isValidate();
        void notifyDirty(T t);
    }
}