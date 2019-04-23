

using System;
using System.Collections.Generic;
using System.Reflection;

public class ModuleNetGameServ : IModule
{
    private SocketTcp mNetGameServ = new SocketTcp();
    private Dictionary<int, Type> id2type = new Dictionary<int, Type>();
    private Dictionary<Type, int> type2id = new Dictionary<Type, int>();
    private MsgRouteMgr routeMgr = new MsgRouteMgr();

    public override bool init()
    {
        foreach (netproto.EMsgIds id in Enum.GetValues(typeof(netproto.EMsgIds)))
        {            
            string tname = id.ToString().Substring(1);
            Type tp = typeof(netproto.EMsgIds).Assembly.GetType("netproto." + tname);
            int iid = (int)id;
            id2type[iid] = tp;
            type2id[tp] = iid;
        }
        id2type[1] = typeof(netproto.PlayerData);
        type2id[typeof(netproto.PlayerData)] = 1;
        id2type[2] = typeof(netproto.PlayerInfo);
        type2id[typeof(netproto.PlayerInfo)] = 2;

        /*检查网络路由消息*/
        foreach (IModule module in ModuleMgr.instance.getAllModules().Values)
        {
            Type moduleType = module.GetType();
            MethodInfo[] allMethod = moduleType.GetMethods(
            BindingFlags.Public | BindingFlags.NonPublic | BindingFlags.Instance);
            foreach (MethodInfo method in allMethod)
            {
                object[] attrs = method.GetCustomAttributes(typeof(MsgRouteAttribute), false);
                if (attrs == null || attrs.Length == 0)
                {
                    continue;
                }

                ParameterInfo[] paramInfos = method.GetParameters();
                if (paramInfos == null || paramInfos.Length != 1)
                {
                    return false;
                }
                ParameterInfo info = paramInfos[0];

                if (!routeMgr.addRoute(type2id[info.ParameterType], module, method))
                {
                    return false;
                }
            }
        }
        
        return true;
    }

    public override bool prestart()
    {
        return true;
    }

    public override void uninit()
    {

    }

    public override void update(int delta)
    {
        mNetGameServ.recv(this.processNetData);
    }

    public int getBytesSend()
    {
        return mNetGameServ.getBytesSend();
    }

    public int getBytesRecv()
    {
        return mNetGameServ.getBytesRecv();
    }

    private byte[] processNetData(byte[] data, int len)
    {
        /*拆包，数据分析*/
        /*
         * 2字节总长度，2字节消息id，消息体
         * */
        int readpos = 0;
        while (len - readpos >= 6)
        {
            int totallen = (int)((uint)data[readpos+0]<<24 | 
                (uint)data[readpos + 1] << 16 |
                (uint)data[readpos + 2] << 8 |
                (uint)data[readpos + 3]);
            if (readpos + totallen > len)
            {
                break;
            }
            int msgid = data[readpos + 4]<<8 | data[readpos + 5];

            byte[] msgbody = new byte[totallen - 6];
            Buffer.BlockCopy(data, readpos + 6 , msgbody, 0, totallen - 6);
            readpos += totallen;

            using (System.IO.MemoryStream mem = new System.IO.MemoryStream(msgbody))
            {
                object obj = ProtoBuf.Serializer.Deserialize(id2type[msgid], mem);                
                routeMgr.getRoute(msgid).call(obj);
            }
        }

        byte[] left = new byte[len-readpos];
        Buffer.BlockCopy(data, readpos, left, 0, len - readpos);

        return left;
    }

    public bool send<T>(T req)
    {
        if(!isNetConnect())
        {
            SysLog.debug("send failed, net not connect");
            return false;
        }

        /*
         * 封包，5字节总长度，2字节消息id，消息体
         * */
        byte[] reqdata = null;
        using (System.IO.MemoryStream mem = new System.IO.MemoryStream())
        {
            ProtoBuf.Serializer.Serialize(mem, req);            
            reqdata = mem.ToArray();
        }

        int msgid = type2id[req.GetType()];


        int totallen = reqdata.Length + 6;

        byte[] buf = new byte[totallen];
        buf[0] = (byte)((0xff000000 & totallen) >> 24);
        buf[1] = (byte)((0x00ff0000 & totallen) >> 16);
        buf[2] = (byte)((0x0000ff00 & totallen) >> 8);
        buf[3] = (byte)(0xff & totallen);

        buf[4] = (byte)((0xff00 & msgid) >> 8); 
        buf[5] = (byte)(0xff & msgid);

        Buffer.BlockCopy(reqdata, 0, buf, 6, reqdata.Length);

        SysLog.debug("send request: " + req.GetType());
        if(!mNetGameServ.send(buf))
        {
            SysLog.debug("send request failed.");
            return false;
        }

        return true;
    }

    public bool isNetConnect()
    {
        return mNetGameServ.isConnect();
    }

    public bool startConnect(string url, int port)
    {
        return mNetGameServ.connect(url, port);
    }

    public void stopConnect()
    {
        mNetGameServ.close();
    }
}
