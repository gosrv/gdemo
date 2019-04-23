

public class TextShow
{
    public static string strcolor(string str, string valuecolor)
    {
        return string.Format("<color=#{0}>{1}</color>", valuecolor, str);
    }
}