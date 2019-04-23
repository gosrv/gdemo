

public abstract class IModule
{
    public abstract bool init();
    public abstract void uninit();
    public abstract bool prestart();
    public abstract void update(int delta);
}
