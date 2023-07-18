basic: #Basic
basic: {
	message: field: "field"
	messageNested: {
		inner: {
			innerEnum:      "Inner_InnerEnum_ONE"
			innerInnerEnum: "Inner_Inner_InnerEnum_ONE"
		}
		innerInner: {
			enum: "Inner_Inner_InnerEnum_ONE"
		}
		innerEnum:           "InnerEnum_ONE"
		innerInnerEnum:      "Inner_InnerEnum_ONE"
		innerInnerInnerEnum: "Inner_Inner_InnerEnum_ONE"
	}
	enum:    "ONE"
	bool:    true
	string:  "string"
	bytes:   'bytes'
	int32:   123
	fixed32: 123
	uint32:  123
	int64:   123
	fixed64: 123
	uint64:  123
	float:   123.456
	double:  123.456
	mapStringString: {
		"key1": "value1"
		"key2": "value2"
	}
	mapStringMessage: {
		"key1": field: "message1"
		"key2": field: "message2"
	}
	mapKeyConvertedToString: {
		"123": "123"
		"456": "456"
	}
	strings: ["one", "two", "three"]
	messages: [{field: "1"}, {field: "2"}]
	oneofString:      "oneOfString"
	int32NonOptional: 123
	messageNonOptional: field: "field"
}

wellKnownTypes: #WellKnownTypes
wellKnownTypes: {
	any: {
		this: "is any value"
	}
	struct: {
		this: "is open struct"
	}
	value: {
		this: "is any value"
	}
	list: ["123", 123, '123']
	"null": null
	bool:   true
	string: "string"
	bytes:  'bytes'
	int32:  123
	uint32: 123
	int64:  123
	uint64: 123
	float:  123.456
	double: 123.456
	empty: {}
	timestamp: "2023-07-17T16:23:00+09:00"
	duration:  "1h2m3s"
	fieldMask: paths: ["a.b.c"]
}
fieldOptions: #FieldOptions
fieldOptions: {
	name: "y"
	age:  20
}
