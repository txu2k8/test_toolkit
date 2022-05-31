package runner

// func assertRunTestCases(t *testing.T) {
// 	testcase1 := &TestCase{
// 		Config: NewConfig("TestCase1").
// 			SetBaseURL("http://httpbin.org"),
// 		TestSteps: []IStep{
// 			NewStep("testcase1-step1").
// 				GET("/headers").
// 				Validate().
// 				AssertEqual("status_code", 200, "check status code").
// 				AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
// 			NewStep("testcase1-step2").
// 				GET("/user-agent").
// 				Validate().
// 				AssertEqual("status_code", 200, "check status code").
// 				AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
// 			NewStep("testcase1-step3").CallRefCase(
// 				&TestCase{
// 					Config: NewConfig("testcase1-step3-ref-case").SetBaseURL("http://httpbin.org"),
// 					TestSteps: []IStep{
// 						NewStep("ip").
// 							GET("/ip").
// 							Validate().
// 							AssertEqual("status_code", 200, "check status code").
// 							AssertEqual("headers.\"Content-Type\"", "application/json", "check http response Content-Type"),
// 					},
// 				},
// 			),
// 			NewStep("testcase1-step4").CallRefCase(&demoTestCaseWithPluginJSONPath),
// 		},
// 	}
// 	testcase2 := &TestCase{
// 		Config: NewConfig("TestCase2").SetWeight(3),
// 	}

// 	r := NewRunner(t)
// 	r.SetPluginLogOn()
// 	err := r.Run(testcase1, testcase2)
// 	if err != nil {
// 		t.Fatalf("run testcase error: %v", err)
// 	}
// }
