package com.example.proxy;

import static org.junit.Assert.assertEquals;

import java.net.URI;

import org.junit.Test;

public class RequestPathTest {

	private RequestPath requestPath;

	@Test
	public void testGetPath() {
		requestPath = new RequestPath(URI.create("https://my-proxy.example.com/path%20to/resource"));

		String result = requestPath.getPath();
		assertEquals("/path to/resource", result);
	}

}