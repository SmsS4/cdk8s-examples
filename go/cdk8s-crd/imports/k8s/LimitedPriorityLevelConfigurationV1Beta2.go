// k8s
package k8s


// LimitedPriorityLevelConfiguration specifies how to handle requests that are subject to limits.
//
// It addresses two issues:
// - How are requests for this priority level limited?
// - What should be done with requests that exceed the limit?
type LimitedPriorityLevelConfigurationV1Beta2 struct {
	// `assuredConcurrencyShares` (ACS) configures the execution limit, which is a limit on the number of requests of this priority level that may be exeucting at a given time.
	//
	// ACS must be a positive number. The server's concurrency limit (SCL) is divided among the concurrency-controlled priority levels in proportion to their assured concurrency shares. This produces the assured concurrency value (ACV) --- the number of requests that may be executing at a time --- for each such priority level:
	//
	// ACV(l) = ceil( SCL * ACS(l) / ( sum[priority levels k] ACS(k) ) )
	//
	// bigger numbers of ACS mean more reserved concurrent requests (at the expense of every other PL). This field has a default value of 30.
	AssuredConcurrencyShares *float64 `field:"optional" json:"assuredConcurrencyShares" yaml:"assuredConcurrencyShares"`
	// `borrowingLimitPercent`, if present, configures a limit on how many seats this priority level can borrow from other priority levels.
	//
	// The limit is known as this level's BorrowingConcurrencyLimit (BorrowingCL) and is a limit on the total number of seats that this level may borrow at any one time. This field holds the ratio of that limit to the level's nominal concurrency limit. When this field is non-nil, it must hold a non-negative integer and the limit is calculated as follows.
	//
	// BorrowingCL(i) = round( NominalCL(i) * borrowingLimitPercent(i)/100.0 )
	//
	// The value of this field can be more than 100, implying that this priority level can borrow a number of seats that is greater than its own nominal concurrency limit (NominalCL). When this field is left `nil`, the limit is effectively infinite.
	BorrowingLimitPercent *float64 `field:"optional" json:"borrowingLimitPercent" yaml:"borrowingLimitPercent"`
	// `lendablePercent` prescribes the fraction of the level's NominalCL that can be borrowed by other priority levels.
	//
	// The value of this field must be between 0 and 100, inclusive, and it defaults to 0. The number of seats that other levels can borrow from this level, known as this level's LendableConcurrencyLimit (LendableCL), is defined as follows.
	//
	// LendableCL(i) = round( NominalCL(i) * lendablePercent(i)/100.0 )
	LendablePercent *float64 `field:"optional" json:"lendablePercent" yaml:"lendablePercent"`
	// `limitResponse` indicates what to do with requests that can not be executed right now.
	LimitResponse *LimitResponseV1Beta2 `field:"optional" json:"limitResponse" yaml:"limitResponse"`
}

