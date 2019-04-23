


public class ToggleEvent
{
    private object mData;
    public delegate void callback(bool value, object data);
    private callback mcb;

    public ToggleEvent(callback cb, object data)
    {
        mData = data;
        mcb = cb;
    }

    public void OnToggleEvent(bool value)
    {
        mcb(value, mData);
    }
}