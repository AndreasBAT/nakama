// Code generated by github.com/gobuffalo/packr. DO NOT EDIT.

package migrate

import "github.com/gobuffalo/packr"

// You can use the "packr clean" command to clean up this,
// and any other packr generated files.
func init() {
	packr.PackJSONBytes("./sql", "20180103142001_initial_schema.sql", "\"LyoNCiAqIENvcHlyaWdodCAyMDE4IFRoZSBOYWthbWEgQXV0aG9ycw0KICoNCiAqIExpY2Vuc2VkIHVuZGVyIHRoZSBBcGFjaGUgTGljZW5zZSwgVmVyc2lvbiAyLjAgKHRoZSAiTGljZW5zZSIpOw0KICogeW91IG1heSBub3QgdXNlIHRoaXMgZmlsZSBleGNlcHQgaW4gY29tcGxpYW5jZSB3aXRoIHRoZSBMaWNlbnNlLg0KICogWW91IG1heSBvYnRhaW4gYSBjb3B5IG9mIHRoZSBMaWNlbnNlIGF0DQogKg0KICogaHR0cDovL3d3dy5hcGFjaGUub3JnL2xpY2Vuc2VzL0xJQ0VOU0UtMi4wDQogKg0KICogVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQ0KICogZGlzdHJpYnV0ZWQgdW5kZXIgdGhlIExpY2Vuc2UgaXMgZGlzdHJpYnV0ZWQgb24gYW4gIkFTIElTIiBCQVNJUywNCiAqIFdJVEhPVVQgV0FSUkFOVElFUyBPUiBDT05ESVRJT05TIE9GIEFOWSBLSU5ELCBlaXRoZXIgZXhwcmVzcyBvciBpbXBsaWVkLg0KICogU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZA0KICogbGltaXRhdGlvbnMgdW5kZXIgdGhlIExpY2Vuc2UuDQogKi8NCg0KLS0gK21pZ3JhdGUgVXANCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIHVzZXJzICgNCiAgICBQUklNQVJZIEtFWSAoaWQpLA0KDQogICAgaWQgICAgICAgICAgICBVVUlEICAgICAgICAgIE5PVCBOVUxMLA0KICAgIHVzZXJuYW1lICAgICAgVkFSQ0hBUigxMjgpICBDT05TVFJBSU5UIHVzZXJzX3VzZXJuYW1lX2tleSBVTklRVUUgTk9UIE5VTEwsDQogICAgZGlzcGxheV9uYW1lICBWQVJDSEFSKDI1NSksDQogICAgYXZhdGFyX3VybCAgICBWQVJDSEFSKDUxMiksDQogICAgLS0gaHR0cHM6Ly90b29scy5pZXRmLm9yZy9odG1sL2JjcDQ3DQogICAgbGFuZ190YWcgICAgICBWQVJDSEFSKDE4KSAgIERFRkFVTFQgJ2VuJywNCiAgICBsb2NhdGlvbiAgICAgIFZBUkNIQVIoMjU1KSwgLS0gZS5nLiAiU2FuIEZyYW5jaXNjbywgQ0EiDQogICAgdGltZXpvbmUgICAgICBWQVJDSEFSKDI1NSksIC0tIGUuZy4gIlBhY2lmaWMgVGltZSAoVVMgJiBDYW5hZGEpIg0KICAgIG1ldGFkYXRhICAgICAgSlNPTkIgICAgICAgICBERUZBVUxUICd7fScgTk9UIE5VTEwsDQogICAgd2FsbGV0ICAgICAgICBKU09OQiAgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwNCiAgICBlbWFpbCAgICAgICAgIFZBUkNIQVIoMjU1KSAgVU5JUVVFLA0KICAgIHBhc3N3b3JkICAgICAgQllURUEgICAgICAgICBDSEVDSyAobGVuZ3RoKHBhc3N3b3JkKSA8IDMyMDAwKSwNCiAgICBmYWNlYm9va19pZCAgIFZBUkNIQVIoMTI4KSAgVU5JUVVFLA0KICAgIGdvb2dsZV9pZCAgICAgVkFSQ0hBUigxMjgpICBVTklRVUUsDQogICAgZ2FtZWNlbnRlcl9pZCBWQVJDSEFSKDEyOCkgIFVOSVFVRSwNCiAgICBzdGVhbV9pZCAgICAgIFZBUkNIQVIoMTI4KSAgVU5JUVVFLA0KICAgIGN1c3RvbV9pZCAgICAgVkFSQ0hBUigxMjgpICBVTklRVUUsDQogICAgZWRnZV9jb3VudCAgICBJTlQgICAgICAgICAgIERFRkFVTFQgMCBDSEVDSyAoZWRnZV9jb3VudCA+PSAwKSBOT1QgTlVMTCwNCiAgICBjcmVhdGVfdGltZSAgIFRJTUVTVEFNUFRaICAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwNCiAgICB1cGRhdGVfdGltZSAgIFRJTUVTVEFNUFRaICAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwNCiAgICB2ZXJpZnlfdGltZSAgIFRJTUVTVEFNUFRaICAgREVGQVVMVCAnMTk3MC0wMS0wMSAwMDowMDowMCBVVEMnIE5PVCBOVUxMLA0KICAgIGRpc2FibGVfdGltZSAgVElNRVNUQU1QVFogICBERUZBVUxUICcxOTcwLTAxLTAxIDAwOjAwOjAwIFVUQycgTk9UIE5VTEwNCik7DQoNCi0tIFNldHVwIFN5c3RlbSB1c2VyLg0KSU5TRVJUIElOVE8gdXNlcnMgKGlkLCB1c2VybmFtZSkNCiAgICBWQUxVRVMgKCcwMDAwMDAwMC0wMDAwLTAwMDAtMDAwMC0wMDAwMDAwMDAwMDAnLCAnJykNCiAgICBPTiBDT05GTElDVChpZCkgRE8gTk9USElORzsNCg0KQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgdXNlcl9kZXZpY2UgKA0KICAgIFBSSU1BUlkgS0VZIChpZCksDQogICAgRk9SRUlHTiBLRVkgKHVzZXJfaWQpIFJFRkVSRU5DRVMgdXNlcnMgKGlkKSBPTiBERUxFVEUgQ0FTQ0FERSwNCg0KICAgIGlkICAgICAgVkFSQ0hBUigxMjgpIE5PVCBOVUxMLA0KICAgIHVzZXJfaWQgVVVJRCAgICAgICAgIE5PVCBOVUxMLA0KDQogICAgVU5JUVVFICh1c2VyX2lkLCBpZCkNCik7DQoNCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIHVzZXJfZWRnZSAoDQogICAgUFJJTUFSWSBLRVkgKHNvdXJjZV9pZCwgc3RhdGUsIHBvc2l0aW9uKSwNCiAgICBGT1JFSUdOIEtFWSAoc291cmNlX2lkKSAgICAgIFJFRkVSRU5DRVMgdXNlcnMgKGlkKSBPTiBERUxFVEUgQ0FTQ0FERSwNCiAgICBGT1JFSUdOIEtFWSAoZGVzdGluYXRpb25faWQpIFJFRkVSRU5DRVMgdXNlcnMgKGlkKSBPTiBERUxFVEUgQ0FTQ0FERSwNCg0KICAgIHNvdXJjZV9pZCAgICAgIFVVSUQgICAgICAgIENIRUNLIChzb3VyY2VfaWQgPD4gJzAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCcpIE5PVCBOVUxMLA0KICAgIHBvc2l0aW9uICAgICAgIEJJR0lOVCAgICAgIE5PVCBOVUxMLCAtLSBVc2VkIGZvciBzb3J0IG9yZGVyIG9uIHJvd3MuDQogICAgdXBkYXRlX3RpbWUgICAgVElNRVNUQU1QVFogREVGQVVMVCBub3coKSBOT1QgTlVMTCwNCiAgICBkZXN0aW5hdGlvbl9pZCBVVUlEICAgICAgICBDSEVDSyAoZGVzdGluYXRpb25faWQgPD4gJzAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCcpIE5PVCBOVUxMLA0KICAgIHN0YXRlICAgICAgICAgIFNNQUxMSU5UICAgIERFRkFVTFQgMCBOT1QgTlVMTCwgLS0gZnJpZW5kKDApLCBpbnZpdGVfc2VudCgxKSwgaW52aXRlX3JlY2VpdmVkKDIpLCBibG9ja2VkKDMpLCBkZWxldGVkKDQpLCBhcmNoaXZlZCg1KQ0KDQogICAgVU5JUVVFIChzb3VyY2VfaWQsIGRlc3RpbmF0aW9uX2lkKQ0KKTsNCkNSRUFURSBJTkRFWCBJRiBOT1QgRVhJU1RTIHVzZXJfZWRnZV9hdXRvX2luZGV4X2ZrX2Rlc3RpbmF0aW9uX2lkX3JlZl91c2VycyBPTiB1c2VyX2VkZ2UgKGRlc3RpbmF0aW9uX2lkKTsNCg0KQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgbm90aWZpY2F0aW9uICgNCiAgICAtLSBGSVhNRTogY29ja3JvYWNoJ3MgYW5hbHlzZXIgaXMgbm90IGNsZXZlciBlbm91Z2ggd2hlbiBjcmVhdGVfdGltZSBoYXMgREVTQyBtb2RlIG9uIHRoZSBpbmRleC4NCiAgICBQUklNQVJZIEtFWSAodXNlcl9pZCwgY3JlYXRlX3RpbWUsIGlkKSwNCiAgICBGT1JFSUdOIEtFWSAodXNlcl9pZCkgUkVGRVJFTkNFUyB1c2VycyAoaWQpIE9OIERFTEVURSBDQVNDQURFLA0KDQogICAgaWQgICAgICAgICAgVVVJRCAgICAgICAgIENPTlNUUkFJTlQgbm90aWZpY2F0aW9uX2lkX2tleSBVTklRVUUgTk9UIE5VTEwsDQogICAgdXNlcl9pZCAgICAgVVVJRCAgICAgICAgIE5PVCBOVUxMLA0KICAgIHN1YmplY3QgICAgIFZBUkNIQVIoMjU1KSBOT1QgTlVMTCwNCiAgICBjb250ZW50ICAgICBKU09OQiAgICAgICAgREVGQVVMVCAne30nIE5PVCBOVUxMLA0KICAgIGNvZGUgICAgICAgIFNNQUxMSU5UICAgICBOT1QgTlVMTCwgLS0gTmVnYXRpdmUgdmFsdWVzIGFyZSBzeXN0ZW0gcmVzZXJ2ZWQuDQogICAgc2VuZGVyX2lkICAgVVVJRCAgICAgICAgIE5PVCBOVUxMLA0KICAgIGNyZWF0ZV90aW1lIFRJTUVTVEFNUFRaICBERUZBVUxUIG5vdygpIE5PVCBOVUxMDQopOw0KDQpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyBzdG9yYWdlICgNCiAgICBQUklNQVJZIEtFWSAoY29sbGVjdGlvbiwgcmVhZCwga2V5LCB1c2VyX2lkKSwNCiAgICBGT1JFSUdOIEtFWSAodXNlcl9pZCkgUkVGRVJFTkNFUyB1c2VycyAoaWQpIE9OIERFTEVURSBDQVNDQURFLA0KDQogICAgY29sbGVjdGlvbiAgVkFSQ0hBUigxMjgpIE5PVCBOVUxMLA0KICAgIGtleSAgICAgICAgIFZBUkNIQVIoMTI4KSBOT1QgTlVMTCwNCiAgICB1c2VyX2lkICAgICBVVUlEICAgICAgICAgTk9UIE5VTEwsDQogICAgdmFsdWUgICAgICAgSlNPTkIgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwNCiAgICB2ZXJzaW9uICAgICBWQVJDSEFSKDMyKSAgTk9UIE5VTEwsIC0tIG1kNSBoYXNoIG9mIHZhbHVlIG9iamVjdC4NCiAgICByZWFkICAgICAgICBTTUFMTElOVCAgICAgREVGQVVMVCAxIENIRUNLIChyZWFkID49IDApIE5PVCBOVUxMLA0KICAgIHdyaXRlICAgICAgIFNNQUxMSU5UICAgICBERUZBVUxUIDEgQ0hFQ0sgKHdyaXRlID49IDApIE5PVCBOVUxMLA0KICAgIGNyZWF0ZV90aW1lIFRJTUVTVEFNUFRaICBERUZBVUxUIG5vdygpIE5PVCBOVUxMLA0KICAgIHVwZGF0ZV90aW1lIFRJTUVTVEFNUFRaICBERUZBVUxUIG5vdygpIE5PVCBOVUxMLA0KDQogICAgVU5JUVVFIChjb2xsZWN0aW9uLCBrZXksIHVzZXJfaWQpDQopOw0KQ1JFQVRFIElOREVYIElGIE5PVCBFWElTVFMgY29sbGVjdGlvbl9yZWFkX3VzZXJfaWRfa2V5X2lkeCBPTiBzdG9yYWdlIChjb2xsZWN0aW9uLCByZWFkLCB1c2VyX2lkLCBrZXkpOw0KQ1JFQVRFIElOREVYIElGIE5PVCBFWElTVFMgc3RvcmFnZV9hdXRvX2luZGV4X2ZrX3VzZXJfaWRfcmVmX3VzZXJzIE9OIHN0b3JhZ2UgKHVzZXJfaWQpOw0KDQpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyBtZXNzYWdlICgNCiAgUFJJTUFSWSBLRVkgKHN0cmVhbV9tb2RlLCBzdHJlYW1fc3ViamVjdCwgc3RyZWFtX2Rlc2NyaXB0b3IsIHN0cmVhbV9sYWJlbCwgY3JlYXRlX3RpbWUsIGlkKSwNCiAgRk9SRUlHTiBLRVkgKHNlbmRlcl9pZCkgUkVGRVJFTkNFUyB1c2VycyAoaWQpIE9OIERFTEVURSBDQVNDQURFLA0KDQogIGlkICAgICAgICAgICAgICAgIFVVSUQgICAgICAgICBVTklRVUUgTk9UIE5VTEwsDQogIC0tIGNoYXQoMCksIGNoYXRfdXBkYXRlKDEpLCBjaGF0X3JlbW92ZSgyKSwgZ3JvdXBfam9pbigzKSwgZ3JvdXBfYWRkKDQpLCBncm91cF9sZWF2ZSg1KSwgZ3JvdXBfa2ljayg2KSwgZ3JvdXBfcHJvbW90ZWQoNykNCiAgY29kZSAgICAgICAgICAgICAgU01BTExJTlQgICAgIERFRkFVTFQgMCBOT1QgTlVMTCwNCiAgc2VuZGVyX2lkICAgICAgICAgVVVJRCAgICAgICAgIE5PVCBOVUxMLA0KICB1c2VybmFtZSAgICAgICAgICBWQVJDSEFSKDEyOCkgTk9UIE5VTEwsDQogIHN0cmVhbV9tb2RlICAgICAgIFNNQUxMSU5UICAgICBOT1QgTlVMTCwNCiAgc3RyZWFtX3N1YmplY3QgICAgVVVJRCAgICAgICAgIE5PVCBOVUxMLA0KICBzdHJlYW1fZGVzY3JpcHRvciBVVUlEICAgICAgICAgTk9UIE5VTEwsDQogIHN0cmVhbV9sYWJlbCAgICAgIFZBUkNIQVIoMTI4KSBOT1QgTlVMTCwNCiAgY29udGVudCAgICAgICAgICAgSlNPTkIgICAgICAgIERFRkFVTFQgJ3t9JyBOT1QgTlVMTCwNCiAgY3JlYXRlX3RpbWUgICAgICAgVElNRVNUQU1QVFogIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsDQogIHVwZGF0ZV90aW1lICAgICAgIFRJTUVTVEFNUFRaICBERUZBVUxUIG5vdygpIE5PVCBOVUxMLA0KDQogIFVOSVFVRSAoc2VuZGVyX2lkLCBpZCkNCik7DQoNCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIGxlYWRlcmJvYXJkICgNCiAgUFJJTUFSWSBLRVkgKGlkKSwNCg0KICBpZCAgICAgICAgICAgICBWQVJDSEFSKDEyOCkgTk9UIE5VTEwsDQogIGF1dGhvcml0YXRpdmUgIEJPT0xFQU4gICAgICBERUZBVUxUIEZBTFNFLA0KICBzb3J0X29yZGVyICAgICBTTUFMTElOVCAgICAgREVGQVVMVCAxIE5PVCBOVUxMLCAtLSBhc2MoMCksIGRlc2MoMSkNCiAgb3BlcmF0b3IgICAgICAgU01BTExJTlQgICAgIERFRkFVTFQgMCBOT1QgTlVMTCwgLS0gYmVzdCgwKSwgc2V0KDEpLCBpbmNyZW1lbnQoMiksIGRlY3JlbWVudCgzKQ0KICByZXNldF9zY2hlZHVsZSBWQVJDSEFSKDY0KSwgLS0gZS5nLiBjcm9uIGZvcm1hdDogIiogKiAqICogKiAqICoiDQogIG1ldGFkYXRhICAgICAgIEpTT05CICAgICAgICBERUZBVUxUICd7fScgTk9UIE5VTEwsDQogIGNyZWF0ZV90aW1lICAgIFRJTUVTVEFNUFRaICBERUZBVUxUIG5vdygpIE5PVCBOVUxMDQopOw0KDQpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyBsZWFkZXJib2FyZF9yZWNvcmQgKA0KICBQUklNQVJZIEtFWSAobGVhZGVyYm9hcmRfaWQsIGV4cGlyeV90aW1lLCBzY29yZSwgc3Vic2NvcmUsIG93bmVyX2lkKSwNCiAgRk9SRUlHTiBLRVkgKGxlYWRlcmJvYXJkX2lkKSBSRUZFUkVOQ0VTIGxlYWRlcmJvYXJkIChpZCkgT04gREVMRVRFIENBU0NBREUsDQoNCiAgbGVhZGVyYm9hcmRfaWQgVkFSQ0hBUigxMjgpICBOT1QgTlVMTCwNCiAgb3duZXJfaWQgICAgICAgVVVJRCAgICAgICAgICBOT1QgTlVMTCwNCiAgdXNlcm5hbWUgICAgICAgVkFSQ0hBUigxMjgpLA0KICBzY29yZSAgICAgICAgICBCSUdJTlQgICAgICAgIERFRkFVTFQgMCBDSEVDSyAoc2NvcmUgPj0gMCkgTk9UIE5VTEwsDQogIHN1YnNjb3JlICAgICAgIEJJR0lOVCAgICAgICAgREVGQVVMVCAwIENIRUNLIChzdWJzY29yZSA+PSAwKSBOT1QgTlVMTCwNCiAgbnVtX3Njb3JlICAgICAgSU5UICAgICAgICAgICBERUZBVUxUIDEgQ0hFQ0sgKG51bV9zY29yZSA+PSAwKSBOT1QgTlVMTCwNCiAgbWV0YWRhdGEgICAgICAgSlNPTkIgICAgICAgICBERUZBVUxUICd7fScgTk9UIE5VTEwsDQogIGNyZWF0ZV90aW1lICAgIFRJTUVTVEFNUFRaICAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwNCiAgdXBkYXRlX3RpbWUgICAgVElNRVNUQU1QVFogICBERUZBVUxUIG5vdygpIE5PVCBOVUxMLA0KICBleHBpcnlfdGltZSAgICBUSU1FU1RBTVBUWiAgIERFRkFVTFQgJzE5NzAtMDEtMDEgMDA6MDA6MDAgVVRDJyBOT1QgTlVMTCwNCg0KICBVTklRVUUgKG93bmVyX2lkLCBsZWFkZXJib2FyZF9pZCwgZXhwaXJ5X3RpbWUpDQopOw0KDQpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyB3YWxsZXRfbGVkZ2VyICgNCiAgUFJJTUFSWSBLRVkgKHVzZXJfaWQsIGNyZWF0ZV90aW1lLCBpZCksDQogIEZPUkVJR04gS0VZICh1c2VyX2lkKSBSRUZFUkVOQ0VTIHVzZXJzIChpZCkgT04gREVMRVRFIENBU0NBREUsDQoNCiAgaWQgICAgICAgICAgVVVJRCAgICAgICAgVU5JUVVFIE5PVCBOVUxMLA0KICB1c2VyX2lkICAgICBVVUlEICAgICAgICBOT1QgTlVMTCwNCiAgY2hhbmdlc2V0ICAgSlNPTkIgICAgICAgTk9UIE5VTEwsDQogIG1ldGFkYXRhICAgIEpTT05CICAgICAgIE5PVCBOVUxMLA0KICBjcmVhdGVfdGltZSBUSU1FU1RBTVBUWiBERUZBVUxUIG5vdygpIE5PVCBOVUxMLA0KICB1cGRhdGVfdGltZSBUSU1FU1RBTVBUWiBERUZBVUxUIG5vdygpIE5PVCBOVUxMDQopOw0KDQpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyB1c2VyX3RvbWJzdG9uZSAoDQogIFBSSU1BUlkgS0VZIChjcmVhdGVfdGltZSwgdXNlcl9pZCksDQoNCiAgdXNlcl9pZCAgICAgICAgVVVJRCAgICAgICAgVU5JUVVFIE5PVCBOVUxMLA0KICBjcmVhdGVfdGltZSAgICBUSU1FU1RBTVBUWiBERUZBVUxUIG5vdygpIE5PVCBOVUxMDQopOw0KDQpDUkVBVEUgVEFCTEUgSUYgTk9UIEVYSVNUUyBncm91cHMgKA0KICBQUklNQVJZIEtFWSAoZGlzYWJsZV90aW1lLCBsYW5nX3RhZywgZWRnZV9jb3VudCwgaWQpLA0KDQogIGlkICAgICAgICAgICBVVUlEICAgICAgICAgIFVOSVFVRSBOT1QgTlVMTCwNCiAgY3JlYXRvcl9pZCAgIFVVSUQgICAgICAgICAgTk9UIE5VTEwsDQogIG5hbWUgICAgICAgICBWQVJDSEFSKDI1NSkgIENPTlNUUkFJTlQgZ3JvdXBzX25hbWVfa2V5IFVOSVFVRSBOT1QgTlVMTCwNCiAgZGVzY3JpcHRpb24gIFZBUkNIQVIoMjU1KSwNCiAgYXZhdGFyX3VybCAgIFZBUkNIQVIoNTEyKSwNCiAgLS0gaHR0cHM6Ly90b29scy5pZXRmLm9yZy9odG1sL2JjcDQ3DQogIGxhbmdfdGFnICAgICBWQVJDSEFSKDE4KSAgIERFRkFVTFQgJ2VuJywNCiAgbWV0YWRhdGEgICAgIEpTT05CICAgICAgICAgREVGQVVMVCAne30nIE5PVCBOVUxMLA0KICBzdGF0ZSAgICAgICAgU01BTExJTlQgICAgICBERUZBVUxUIDAgQ0hFQ0sgKHN0YXRlID49IDApIE5PVCBOVUxMLCAtLSBvcGVuKDApLCBjbG9zZWQoMSkNCiAgZWRnZV9jb3VudCAgIElOVCAgICAgICAgICAgREVGQVVMVCAwIENIRUNLIChlZGdlX2NvdW50ID49IDEgQU5EIGVkZ2VfY291bnQgPD0gbWF4X2NvdW50KSBOT1QgTlVMTCwNCiAgbWF4X2NvdW50ICAgIElOVCAgICAgICAgICAgREVGQVVMVCAxMDAgQ0hFQ0sgKG1heF9jb3VudCA+PSAxKSBOT1QgTlVMTCwNCiAgY3JlYXRlX3RpbWUgIFRJTUVTVEFNUFRaICAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwNCiAgdXBkYXRlX3RpbWUgIFRJTUVTVEFNUFRaICAgREVGQVVMVCBub3coKSBOT1QgTlVMTCwNCiAgZGlzYWJsZV90aW1lIFRJTUVTVEFNUFRaICAgREVGQVVMVCAnMTk3MC0wMS0wMSAwMDowMDowMCBVVEMnIE5PVCBOVUxMDQopOw0KQ1JFQVRFIElOREVYIElGIE5PVCBFWElTVFMgZWRnZV9jb3VudF91cGRhdGVfdGltZV9pZF9pZHggT04gZ3JvdXBzIChkaXNhYmxlX3RpbWUsIGVkZ2VfY291bnQsIHVwZGF0ZV90aW1lLCBpZCk7DQpDUkVBVEUgSU5ERVggSUYgTk9UIEVYSVNUUyB1cGRhdGVfdGltZV9lZGdlX2NvdW50X2lkX2lkeCBPTiBncm91cHMgKGRpc2FibGVfdGltZSwgdXBkYXRlX3RpbWUsIGVkZ2VfY291bnQsIGlkKTsNCg0KQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgZ3JvdXBfZWRnZSAoDQogIFBSSU1BUlkgS0VZIChzb3VyY2VfaWQsIHN0YXRlLCBwb3NpdGlvbiksDQoNCiAgc291cmNlX2lkICAgICAgVVVJRCAgICAgICAgQ0hFQ0sgKHNvdXJjZV9pZCA8PiAnMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwJykgTk9UIE5VTEwsDQogIHBvc2l0aW9uICAgICAgIEJJR0lOVCAgICAgIE5PVCBOVUxMLCAtLSBVc2VkIGZvciBzb3J0IG9yZGVyIG9uIHJvd3MuDQogIHVwZGF0ZV90aW1lICAgIFRJTUVTVEFNUFRaIERFRkFVTFQgbm93KCkgTk9UIE5VTEwsDQogIGRlc3RpbmF0aW9uX2lkIFVVSUQgICAgICAgIENIRUNLIChkZXN0aW5hdGlvbl9pZCA8PiAnMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwJykgTk9UIE5VTEwsDQogIHN0YXRlICAgICAgICAgIFNNQUxMSU5UICAgIERFRkFVTFQgMCBOT1QgTlVMTCwgLS0gc3VwZXJhZG1pbigwKSwgYWRtaW4oMSksIG1lbWJlcigyKSwgam9pbl9yZXF1ZXN0KDMpLCBhcmNoaXZlZCg0KQ0KDQogIFVOSVFVRSAoc291cmNlX2lkLCBkZXN0aW5hdGlvbl9pZCkNCik7DQoNCi0tICttaWdyYXRlIERvd24NCkRST1AgVEFCTEUgSUYgRVhJU1RTIGdyb3VwX2VkZ2U7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyBncm91cHM7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyB1c2VyX3RvbWJzdG9uZTsNCkRST1AgVEFCTEUgSUYgRVhJU1RTIHdhbGxldF9sZWRnZXI7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyBsZWFkZXJib2FyZF9yZWNvcmQ7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyBsZWFkZXJib2FyZDsNCkRST1AgVEFCTEUgSUYgRVhJU1RTIG1lc3NhZ2U7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyBzdG9yYWdlOw0KRFJPUCBUQUJMRSBJRiBFWElTVFMgbm90aWZpY2F0aW9uOw0KRFJPUCBUQUJMRSBJRiBFWElTVFMgdXNlcl9lZGdlOw0KRFJPUCBUQUJMRSBJRiBFWElTVFMgdXNlcl9kZXZpY2U7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyB1c2VyczsNCg==\"")
	packr.PackJSONBytes("./sql", "20180805174141-tournaments.sql", "\"LyoNCiAqIENvcHlyaWdodCAyMDE4IFRoZSBOYWthbWEgQXV0aG9ycw0KICoNCiAqIExpY2Vuc2VkIHVuZGVyIHRoZSBBcGFjaGUgTGljZW5zZSwgVmVyc2lvbiAyLjAgKHRoZSAiTGljZW5zZSIpOw0KICogeW91IG1heSBub3QgdXNlIHRoaXMgZmlsZSBleGNlcHQgaW4gY29tcGxpYW5jZSB3aXRoIHRoZSBMaWNlbnNlLg0KICogWW91IG1heSBvYnRhaW4gYSBjb3B5IG9mIHRoZSBMaWNlbnNlIGF0DQogKg0KICogaHR0cDovL3d3dy5hcGFjaGUub3JnL2xpY2Vuc2VzL0xJQ0VOU0UtMi4wDQogKg0KICogVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQ0KICogZGlzdHJpYnV0ZWQgdW5kZXIgdGhlIExpY2Vuc2UgaXMgZGlzdHJpYnV0ZWQgb24gYW4gIkFTIElTIiBCQVNJUywNCiAqIFdJVEhPVVQgV0FSUkFOVElFUyBPUiBDT05ESVRJT05TIE9GIEFOWSBLSU5ELCBlaXRoZXIgZXhwcmVzcyBvciBpbXBsaWVkLg0KICogU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZA0KICogbGltaXRhdGlvbnMgdW5kZXIgdGhlIExpY2Vuc2UuDQogKi8NCg0KLS0gTk9URTogVGhpcyBtaWdyYXRpb24gbWFudWFsbHkgY29tbWl0cyBpbiBzZXBhcmF0ZSB0cmFuc2FjdGlvbnMgdG8gZW5zdXJlDQotLSB0aGUgc2NoZW1hIHVwZGF0ZXMgYXJlIHNlcXVlbmNlZCBiZWNhdXNlIGNvY2tyb2FjaGRiIGRvZXMgbm90IHN1cHBvcnQNCi0tIGFkZGluZyBDSEVDSyBjb25zdHJhaW50cyB2aWEgIkFMVEVSIFRBQkxFIC4uLiBBREQgQ09MVU1OIiBzdGF0ZW1lbnRzLg0KDQotLSArbWlncmF0ZSBVcCBub3RyYW5zYWN0aW9uDQpCRUdJTjsNCkFMVEVSIFRBQkxFIGxlYWRlcmJvYXJkDQogIEFERCBDT0xVTU4gY2F0ZWdvcnkgICAgICBTTUFMTElOVCAgICAgREVGQVVMVCAwIE5PVCBOVUxMLA0KICBBREQgQ09MVU1OIGRlc2NyaXB0aW9uICAgVkFSQ0hBUigyNTUpIERFRkFVTFQgJycgTk9UIE5VTEwsDQogIEFERCBDT0xVTU4gZHVyYXRpb24gICAgICBJTlQgICAgICAgICAgREVGQVVMVCAwIE5PVCBOVUxMLCAtLSBpbiBzZWNvbmRzLg0KICBBREQgQ09MVU1OIGVuZF90aW1lICAgICAgVElNRVNUQU1QVFogIERFRkFVTFQgJzE5NzAtMDEtMDEgMDA6MDA6MDAgVVRDJyBOT1QgTlVMTCwNCiAgQUREIENPTFVNTiBqb2luX3JlcXVpcmVkIEJPT0xFQU4gICAgICBERUZBVUxUIEZBTFNFIE5PVCBOVUxMLA0KICBBREQgQ09MVU1OIG1heF9zaXplICAgICAgSU5UICAgICAgICAgIERFRkFVTFQgMTAwMDAwMDAwIE5PVCBOVUxMLA0KICBBREQgQ09MVU1OIG1heF9udW1fc2NvcmUgSU5UICAgICAgICAgIERFRkFVTFQgMTAwMDAwMCBOT1QgTlVMTCwgLS0gbWF4IGFsbG93ZWQgc2NvcmUgYXR0ZW1wdHMuDQogIEFERCBDT0xVTU4gdGl0bGUgICAgICAgICBWQVJDSEFSKDI1NSkgREVGQVVMVCAnJyBOT1QgTlVMTCwNCiAgQUREIENPTFVNTiBzaXplICAgICAgICAgIElOVCAgICAgICAgICBERUZBVUxUIDAgTk9UIE5VTEwsDQogIEFERCBDT0xVTU4gc3RhcnRfdGltZSAgICBUSU1FU1RBTVBUWiAgREVGQVVMVCBub3coKSBOT1QgTlVMTDsNCg0KQUxURVIgVEFCTEUgbGVhZGVyYm9hcmRfcmVjb3JkDQogIEFERCBDT0xVTU4gbWF4X251bV9zY29yZSBJTlQgREVGQVVMVCAxMDAwMDAwIE5PVCBOVUxMOw0KQ09NTUlUOw0KDQpCRUdJTjsNCkFMVEVSIFRBQkxFIGxlYWRlcmJvYXJkDQogIEFERCBDT05TVFJBSU5UIGNoZWNrX2NhdGVnb3J5IENIRUNLIChjYXRlZ29yeSA+PSAwKSwNCiAgQUREIENPTlNUUkFJTlQgY2hlY2tfZHVyYXRpb24gQ0hFQ0sgKGR1cmF0aW9uID49IDApLA0KICBBREQgQ09OU1RSQUlOVCBjaGVja19tYXhfc2l6ZSBDSEVDSyAobWF4X3NpemUgPiAwKSwNCiAgQUREIENPTlNUUkFJTlQgY2hlY2tfbWF4X251bV9zY29yZSBDSEVDSyAobWF4X251bV9zY29yZSA+IDApOw0KDQpDUkVBVEUgSU5ERVggSUYgTk9UIEVYSVNUUyBkdXJhdGlvbl9zdGFydF90aW1lX2VuZF90aW1lX2NhdGVnb3J5X2lkeCBPTiBsZWFkZXJib2FyZCAoZHVyYXRpb24sIHN0YXJ0X3RpbWUsIGVuZF90aW1lIERFU0MsIGNhdGVnb3J5KTsNCg0KQ1JFQVRFIElOREVYIElGIE5PVCBFWElTVFMgb3duZXJfaWRfZXhwaXJ5X3RpbWVfbGVhZGVyYm9hcmRfaWRfaWR4IE9OIGxlYWRlcmJvYXJkX3JlY29yZCAob3duZXJfaWQsIGV4cGlyeV90aW1lLCBsZWFkZXJib2FyZF9pZCk7DQoNCkFMVEVSIFRBQkxFIGxlYWRlcmJvYXJkX3JlY29yZA0KICBBREQgQ09OU1RSQUlOVCBjaGVja19tYXhfbnVtX3Njb3JlIENIRUNLIChtYXhfbnVtX3Njb3JlID4gMCk7DQpDT01NSVQ7DQoNCkFMVEVSIFRBQkxFIGxlYWRlcmJvYXJkDQogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfY2F0ZWdvcnksDQogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfZHVyYXRpb24sDQogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfbWF4X3NpemUsDQogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfbWF4X251bV9zY29yZTsNCg0KQUxURVIgVEFCTEUgbGVhZGVyYm9hcmRfcmVjb3JkDQogIFZBTElEQVRFIENPTlNUUkFJTlQgY2hlY2tfbWF4X251bV9zY29yZTsNCg0KLS0gK21pZ3JhdGUgRG93bg0KQUxURVIgVEFCTEUgSUYgRVhJU1RTIGxlYWRlcmJvYXJkDQogIERST1AgQ09MVU1OIElGIEVYSVNUUyBjYXRlZ29yeSwNCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIGRlc2NyaXB0aW9uLA0KICBEUk9QIENPTFVNTiBJRiBFWElTVFMgZHVyYXRpb24sDQogIERST1AgQ09MVU1OIElGIEVYSVNUUyBlbmRfdGltZSwNCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIGpvaW5fcmVxdWlyZWQsDQogIERST1AgQ09MVU1OIElGIEVYSVNUUyBtYXhfc2l6ZSwNCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIG1heF9udW1fc2NvcmUsDQogIERST1AgQ09MVU1OIElGIEVYSVNUUyB0aXRsZSwNCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIHNpemUsDQogIERST1AgQ09MVU1OIElGIEVYSVNUUyBzdGFydF90aW1lOw0KDQpEUk9QIElOREVYIElGIEVYSVNUUyBvd25lcl9pZF9leHBpcnlfdGltZV9sZWFkZXJib2FyZF9pZF9pZHg7DQoNCkFMVEVSIFRBQkxFIElGIEVYSVNUUyBsZWFkZXJib2FyZF9yZWNvcmQNCiAgRFJPUCBDT0xVTU4gSUYgRVhJU1RTIG1heF9udW1fc2NvcmU7DQo=\"")
	packr.PackJSONBytes("./sql", "20190621000000_achievements.sql", "\"LyoNCiAqIENvcHlyaWdodCAyMDE4IFRoZSBOYWthbWEgQXV0aG9ycw0KICoNCiAqIExpY2Vuc2VkIHVuZGVyIHRoZSBBcGFjaGUgTGljZW5zZSwgVmVyc2lvbiAyLjAgKHRoZSAiTGljZW5zZSIpOw0KICogeW91IG1heSBub3QgdXNlIHRoaXMgZmlsZSBleGNlcHQgaW4gY29tcGxpYW5jZSB3aXRoIHRoZSBMaWNlbnNlLg0KICogWW91IG1heSBvYnRhaW4gYSBjb3B5IG9mIHRoZSBMaWNlbnNlIGF0DQogKg0KICogaHR0cDovL3d3dy5hcGFjaGUub3JnL2xpY2Vuc2VzL0xJQ0VOU0UtMi4wDQogKg0KICogVW5sZXNzIHJlcXVpcmVkIGJ5IGFwcGxpY2FibGUgbGF3IG9yIGFncmVlZCB0byBpbiB3cml0aW5nLCBzb2Z0d2FyZQ0KICogZGlzdHJpYnV0ZWQgdW5kZXIgdGhlIExpY2Vuc2UgaXMgZGlzdHJpYnV0ZWQgb24gYW4gIkFTIElTIiBCQVNJUywNCiAqIFdJVEhPVVQgV0FSUkFOVElFUyBPUiBDT05ESVRJT05TIE9GIEFOWSBLSU5ELCBlaXRoZXIgZXhwcmVzcyBvciBpbXBsaWVkLg0KICogU2VlIHRoZSBMaWNlbnNlIGZvciB0aGUgc3BlY2lmaWMgbGFuZ3VhZ2UgZ292ZXJuaW5nIHBlcm1pc3Npb25zIGFuZA0KICogbGltaXRhdGlvbnMgdW5kZXIgdGhlIExpY2Vuc2UuDQogKi8NCg0KLS0gK21pZ3JhdGUgVXANCkNSRUFURSBUQUJMRSBJRiBOT1QgRVhJU1RTIGFjaGlldmVtZW50cyAoDQogICAgaWQgdXVpZCBOT1QgTlVMTCwNCiAgICBuYW1lIHRleHQgTk9UIE5VTEwsDQogICAgZGVzY3JpcHRpb24gdGV4dCBOVUxMLA0KICAgIGluaXRpYWxfc3RhdGUgaW50OCBOT1QgTlVMTCwNCiAgICAidHlwZSIgaW50OCBOT1QgTlVMTCwNCiAgICByZXBlYXRhYmlsaXR5IGludDggTk9UIE5VTEwsDQogICAgdGFyZ2V0X3ZhbHVlIGludDggTlVMTCwNCiAgICBsb2NrZWRfaW1hZ2VfdXJsIHRleHQgTlVMTCwNCiAgICB1bmxvY2tlZF9pbWFnZV91cmwgdGV4dCBOVUxMLA0KICAgIGF1eGlsaWFyeV9kYXRhIGpzb25iIE5VTEwsDQogICAgQ09OU1RSQUlOVCBhY2hpZXZlbWVudHNfcGsgUFJJTUFSWSBLRVkgKGlkKQ0KKTsNCg0KQ1JFQVRFIFRBQkxFIElGIE5PVCBFWElTVFMgYWNoaWV2ZW1lbnRfcHJvZ3Jlc3MgKA0KICAgIGlkIHV1aWQgTk9UIE5VTEwsDQogICAgYWNoaWV2ZW1lbnRfaWQgdXVpZCBOT1QgTlVMTCwNCiAgICB1c2VyX2lkIHV1aWQgTk9UIE5VTEwsDQogICAgYWNoaWV2ZW1lbnRfc3RhdGUgaW50OCBOT1QgTlVMTCwNCiAgICBwcm9ncmVzcyBpbnQ4IE5VTEwsDQogICAgYXdhcmRlZF9hdCB0aW1lc3RhbXB0eiBOVUxMLA0KICAgIGF1eGlsaWFyeV9kYXRhIGpzb25iIE5VTEwsDQogICAgQ09OU1RSQUlOVCBhY2hpZXZlbWVudF9pZF9mayBGT1JFSUdOIEtFWSAoYWNoaWV2ZW1lbnRfaWQpIFJFRkVSRU5DRVMgYWNoaWV2ZW1lbnRzKGlkKSwNCiAgICBDT05TVFJBSU5UIGFjaGlldmVtZW50X293bmVyX2ZrIEZPUkVJR04gS0VZICh1c2VyX2lkKSBSRUZFUkVOQ0VTICJ1c2VycyIoaWQpLA0KICAgIENPTlNUUkFJTlQgYWNoaWV2ZW1lbnRfcHJvZ3Jlc3NfcGsgUFJJTUFSWSBLRVkgKGlkKQ0KKTsNCg0KLS0gK21pZ3JhdGUgRG93bg0KRFJPUCBUQUJMRSBJRiBFWElTVFMgYWNoaWV2ZW1lbnRfcHJvZ3Jlc3M7DQpEUk9QIFRBQkxFIElGIEVYSVNUUyBhY2hpZXZlbWVudHM7\"")
}
