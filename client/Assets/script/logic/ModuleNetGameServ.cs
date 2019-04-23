

using System;

public class ModuleNetGameServ : IModule
{
    private SocketTcp mNetGameServ = new SocketTcp();

    public override bool init()
    {
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
        while (len - readpos >= 4)
        {
            int totallen = data[readpos+0] | data[readpos + 1] << 8;
            if (readpos + totallen > len)
            {
                break;
            }
            int msgid = data[readpos + 2] | data[readpos + 3] << 8;

            byte[] msgbody = new byte[totallen - 4 + 1];
            Buffer.BlockCopy(data, readpos + 4 , msgbody, 0, totallen - 4);
            RouteMethod route = MsgRouteMgr.instance.getRoute(msgid);
            if (route == null)
            {
                SysLog.error("can not get route " + msgid);
                return null;
            }

            route.call(msgbody);

            readpos += totallen;
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
         * 封包，2字节总长度，2字节消息id，消息体
         * */
        /*byte[] reqdata = null;
        using (System.IO.MemoryStream mem = new System.IO.MemoryStream())
        {
            ProtoBuf.Serializer.Serialize(mem, req);            
            reqdata = mem.ToArray();
        }

        string msgname = req.GetType().Name;
        short msgid = (short)(csmsg_ids.getId(msgname));

        if (reqdata.Length + 4 > 65536)
        {
            return false;
        }

        short totallen = (short)(reqdata.Length + 4);

        byte[] buf = new byte[totallen];
        buf[0] = (byte)(0xff & totallen);
        buf[1] = (byte)((0xff00 & totallen) >> 8);
        buf[2] = (byte)(0xff & msgid);
        buf[3] = (byte)((0xff00 & msgid) >> 8);

        Buffer.BlockCopy(reqdata, 0, buf, 4, reqdata.Length);

        SysLog.debug("send request: " + msgname);
        if(!mNetGameServ.send(buf))
        {
            SysLog.debug("send request failed.");
            return false;
        }*/

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
