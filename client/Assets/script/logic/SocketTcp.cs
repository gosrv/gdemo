
using System;
using System.Collections;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using System.Collections.Generic;


public class SocketTcp
{
    private Socket mSocket = null;
    private const int mBufSize = 16384;
    private int mLeftDataPos = 0;
    private byte[] mBuffLeft = new byte[mBufSize*2];
    private byte[] mBuffRecv = new byte[mBufSize];
    private int mBytesSend = 0;
    private int mBytesRecv = 0;

    public int getBytesSend()
    {
        return mBytesSend;
    }

    public int getBytesRecv()
    {
        return mBytesRecv;
    }

    public bool connect(string url, int port)
    {
        close();

        IPAddress addr = IPAddress.Parse(url);
        EndPoint ep = new IPEndPoint(addr, port);
        mSocket = new Socket(addr.AddressFamily, SocketType.Stream, ProtocolType.Tcp);
        mSocket.Connect(ep);
        if (!mSocket.Connected)
        {
            close();
            return false;
        }

        mSocket.Blocking = false;
        return true;
    }

    public delegate byte[] DataCallback(byte[] data, int len);
    public void recv(DataCallback cb)
    {
        if (mSocket == null)
        {
            return;
        }

        if (!mSocket.Poll(0, SelectMode.SelectRead))
        {
            return;
        }

        int rcvlen = 0;
        try
        {
            rcvlen = mSocket.Receive(mBuffRecv);
        }
        catch (System.Exception)
        {            
            rcvlen = -1;
        }
        
        if (rcvlen <= 0)
        {
            close();
            return;
        }
        else
        {
            mBytesRecv += rcvlen;
        }

        Buffer.BlockCopy(mBuffRecv, 0, mBuffLeft, mLeftDataPos, rcvlen);
        mLeftDataPos += rcvlen;
        
        byte[] left = cb(mBuffLeft, mLeftDataPos);
        mLeftDataPos = 0;
        if (left != null)
        {
            Buffer.BlockCopy(left, 0, mBuffLeft, 0, left.Length);
            mLeftDataPos = left.Length;
        }        
    }

    public bool send(byte[] data)
    {
        if (mSocket == null)
        {
            return false;
        }
        if(mSocket.Send(data) != data.Length)
        {
            close();
            return false;
        }
        else
        {
            mBytesSend += data.Length;
        }

        return true;
    }

    public void close()
    {
        if (mSocket != null)
        {
            mSocket.Close();
            mSocket = null;
        }
        mLeftDataPos = 0;
    }

    public bool isConnect()
    {
        return mSocket != null && mSocket.Connected;
    }
}