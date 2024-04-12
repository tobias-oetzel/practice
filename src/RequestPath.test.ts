import { strictEqual } from "assert";
import { RequestPath } from "./RequestPath";

describe("Request Path", function () {
  it("gets a path", function () {
        // Arrange
		const url = new URL("https://my-proxy.example.com/path%20to/resource")

		// Act
		const requestPath = new RequestPath(url)
		const result = requestPath.getPath()

		// Assert
		strictEqual(result, "/path to/resource")
  });
});
