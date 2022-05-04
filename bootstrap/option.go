/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package bootstrap

//type optionFunc func(*loaderConf)
//
//func (fn optionFunc) apply(vc *loaderConf) {
//	fn(vc)
//}

//type Option interface {
//	apply(vc *loaderConf)
//}

type Option struct {
	// loaderConf file extension default yaml
	suffix string

	// loaderConf file path default ./configs
	path string

	// config file application
	name string
}

func New() *Option {
	return &Option{
		suffix: "yaml",
		path:   "./configs",
		name:   "application",
	}
}

func (o *Option) SetName(name string) *Option {
	o.name = name
	return o
}

func (o *Option) SetSuffix(suffix string) *Option {
	o.suffix = suffix
	return o
}

func (o *Option) SetPath(path string) *Option {
	o.path = path
	return o
}

//
//// WithPath set load config path
//func WithPath(path string) Option {
//	return optionFunc(func(conf *loaderConf) {
//		conf.path = path
//	})
//}
//
//func WithName(name string) Option {
//	return optionFunc(func(conf *loaderConf) {
//		conf.name = name
//	})
//}
//
//func WithSuffix(suffix string) Option {
//	return optionFunc(func(conf *loaderConf) {
//		conf.suffix = suffix
//	})
//}