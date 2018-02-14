package dispatcher

import "testing"

func TestGenerateFromConfig(t *testing.T) {
	t.Run("SetInvalidValue", func(t *testing.T) {
		type test struct {
			config    Config
			wantedErr error
		}

		testCases := []test{
			{
				// A pattern that is normally generated
				config: Config{
					MaxWorkers: 2,
					MaxQueues:  10,
					MaxRetry:   0,
				},
				wantedErr: nil,
			},
			{
				// A pattern that is normally generated 2
				config: Config{
					MaxWorkers: 3,
					MaxQueues:  1,
					MaxRetry:   0,
				},
				wantedErr: nil,
			},
			{
				// MaxWorkers wrong pattern
				config: Config{
					MaxWorkers: -1,
					MaxQueues:  1,
					MaxRetry:   0,
				},
				wantedErr: ErrInvalidSetting,
			},
			{
				// MaxQueues is zero pattern
				config: Config{
					MaxWorkers: 1,
					MaxQueues:  0,
					MaxRetry:   0,
				},
				wantedErr: nil,
			},
			{
				// MaxQueues is negative pattern
				config: Config{
					MaxWorkers: 1,
					MaxQueues:  -1,
					MaxRetry:   0,
				},
				wantedErr: ErrInvalidSetting,
			},
			{
				// MaxRetry has an invalid pattern
				config: Config{
					MaxWorkers: 2,
					MaxQueues:  10,
					MaxRetry:   -1,
				},
				wantedErr: ErrInvalidSetting,
			},
		}

		for _, testCase := range testCases {
			if _, err := GenerateFromConfig(testCase.config); err != testCase.wantedErr {
				t.Error("GenerateFromConfig SetInvalidValueTest wanted: ", testCase.wantedErr, "but got", err)
				t.Log(testCase.config)
			}
		}
	})

	t.Run("HasLogger", func(t *testing.T) {
		type test struct {
			config          Config
			wantedHasLogger bool
		}

		testCases := []test{
			{
				config: Config{
					MaxWorkers: 2,
					MaxQueues:  10,
					MaxRetry:   0,

					// set nil
					Logger: nil,
				},
				wantedHasLogger: false,
			},
			{
				config: Config{
					MaxWorkers: 2,
					MaxQueues:  10,
					MaxRetry:   0,

					// set pointer of testLogger logger
					Logger: &testLogger{},
				},
				wantedHasLogger: true,
			},
		}

		for _, testCase := range testCases {
			d, _ := GenerateFromConfig(testCase.config)
			disp := d.(*dispatcher)
			if disp.hasLogger != testCase.wantedHasLogger {
				t.Error("GenerateFromConfig SetInvalidValueTest wanted: ", testCase.wantedHasLogger)
			}
		}
	})

}
