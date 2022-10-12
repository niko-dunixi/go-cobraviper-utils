// Code generated with `make generate`; DO NOT EDIT.
//go:build !generating

package cobraviperutils

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindFlag(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int](
				flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				int(1234),
				GetFlag[int]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "1337")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				int(1337),
				GetFlag[int]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "1337"})
			require.NoError(t, err)
			assert.Equal(t,
				int(1337),
				GetFlag[int]("my-key"),
			)
		})
	})
	t.Run("int32", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int32](
				flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				int32(1234),
				GetFlag[int32]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "1337")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int32](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				int32(1337),
				GetFlag[int32]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int32](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "1337"})
			require.NoError(t, err)
			assert.Equal(t,
				int32(1337),
				GetFlag[int32]("my-key"),
			)
		})
	})
	t.Run("int64", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int64](
				flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				int64(1234),
				GetFlag[int64]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "1337")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int64](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				int64(1337),
				GetFlag[int64]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[int64](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "1337"})
			require.NoError(t, err)
			assert.Equal(t,
				int64(1337),
				GetFlag[int64]("my-key"),
			)
		})
	})
	t.Run("string", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[string](
				flags,
				"my-key",
				"I am the default value",
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				string("I am the default value"),
				GetFlag[string]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "Hey, I'm Niko Dunixi!")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[string](flags,
				"my-key",
				"I am the default value",
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				string("Hey, I'm Niko Dunixi!"),
				GetFlag[string]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[string](flags,
				"my-key",
				"I am the default value",
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "Hey, I'm Niko Dunixi!"})
			require.NoError(t, err)
			assert.Equal(t,
				string("Hey, I'm Niko Dunixi!"),
				GetFlag[string]("my-key"),
			)
		})
	})
	t.Run("uint", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint](
				flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				uint(1234),
				GetFlag[uint]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "1337")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				uint(1337),
				GetFlag[uint]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "1337"})
			require.NoError(t, err)
			assert.Equal(t,
				uint(1337),
				GetFlag[uint]("my-key"),
			)
		})
	})
	t.Run("uint32", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint32](
				flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				uint32(1234),
				GetFlag[uint32]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "1337")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint32](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				uint32(1337),
				GetFlag[uint32]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint32](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "1337"})
			require.NoError(t, err)
			assert.Equal(t,
				uint32(1337),
				GetFlag[uint32]("my-key"),
			)
		})
	})
	t.Run("uint64", func(t *testing.T) {
		// No parameters set by user
		t.Run("default fallback value", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint64](
				flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				uint64(1234),
				GetFlag[uint64]("my-key"),
			)
		})
		// Environment variable set
		t.Run("environment variable", func(t *testing.T) {
			os.Setenv("MY_KEY", "1337")
			defer os.Unsetenv("MY_KEY")

			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint64](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{})
			require.NoError(t, err)
			assert.Equal(t,
				uint64(1337),
				GetFlag[uint64]("my-key"),
			)
		})
		// Flag value set
		t.Run("cli parameter", func(t *testing.T) {
			flags := pflag.NewFlagSet("testing", pflag.ContinueOnError)
			BindFlag[uint64](flags,
				"my-key",
				1234,
				"flag usage or description",
			)

			err := flags.Parse([]string{"--my-key", "1337"})
			require.NoError(t, err)
			assert.Equal(t,
				uint64(1337),
				GetFlag[uint64]("my-key"),
			)
		})
	})
}
