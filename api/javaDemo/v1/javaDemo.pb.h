// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: javaDemo/v1/javaDemo.proto
// Protobuf C++ Version: 5.27.0

#ifndef GOOGLE_PROTOBUF_INCLUDED_javaDemo_2fv1_2fjavaDemo_2eproto_2epb_2eh
#define GOOGLE_PROTOBUF_INCLUDED_javaDemo_2fv1_2fjavaDemo_2eproto_2epb_2eh

#include <limits>
#include <string>
#include <type_traits>
#include <utility>

#include "google/protobuf/runtime_version.h"
#if PROTOBUF_VERSION != 5027000
#error "Protobuf C++ gencode is built with an incompatible version of"
#error "Protobuf C++ headers/runtime. See"
#error "https://protobuf.dev/support/cross-version-runtime-guarantee/#cpp"
#endif
#include "google/protobuf/io/coded_stream.h"
#include "google/protobuf/arena.h"
#include "google/protobuf/arenastring.h"
#include "google/protobuf/generated_message_tctable_decl.h"
#include "google/protobuf/generated_message_util.h"
#include "google/protobuf/metadata_lite.h"
#include "google/protobuf/generated_message_reflection.h"
#include "google/protobuf/message.h"
#include "google/protobuf/repeated_field.h"  // IWYU pragma: export
#include "google/protobuf/extension_set.h"  // IWYU pragma: export
#include "google/protobuf/unknown_field_set.h"
#include "google/api/annotations.pb.h"
// @@protoc_insertion_point(includes)

// Must be included last.
#include "google/protobuf/port_def.inc"

#define PROTOBUF_INTERNAL_EXPORT_javaDemo_2fv1_2fjavaDemo_2eproto

namespace google {
namespace protobuf {
namespace internal {
class AnyMetadata;
}  // namespace internal
}  // namespace protobuf
}  // namespace google

// Internal implementation detail -- do not use these members.
struct TableStruct_javaDemo_2fv1_2fjavaDemo_2eproto {
  static const ::uint32_t offsets[];
};
extern const ::google::protobuf::internal::DescriptorTable
    descriptor_table_javaDemo_2fv1_2fjavaDemo_2eproto;
namespace javaDemo {
namespace v1 {
class HelloRequest;
struct HelloRequestDefaultTypeInternal;
extern HelloRequestDefaultTypeInternal _HelloRequest_default_instance_;
class HelloResponse;
struct HelloResponseDefaultTypeInternal;
extern HelloResponseDefaultTypeInternal _HelloResponse_default_instance_;
}  // namespace v1
}  // namespace javaDemo
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

namespace javaDemo {
namespace v1 {

// ===================================================================


// -------------------------------------------------------------------

class HelloResponse final : public ::google::protobuf::Message
/* @@protoc_insertion_point(class_definition:javaDemo.v1.HelloResponse) */ {
 public:
  inline HelloResponse() : HelloResponse(nullptr) {}
  ~HelloResponse() override;
  template <typename = void>
  explicit PROTOBUF_CONSTEXPR HelloResponse(
      ::google::protobuf::internal::ConstantInitialized);

  inline HelloResponse(const HelloResponse& from) : HelloResponse(nullptr, from) {}
  inline HelloResponse(HelloResponse&& from) noexcept
      : HelloResponse(nullptr, std::move(from)) {}
  inline HelloResponse& operator=(const HelloResponse& from) {
    CopyFrom(from);
    return *this;
  }
  inline HelloResponse& operator=(HelloResponse&& from) noexcept {
    if (this == &from) return *this;
    if (GetArena() == from.GetArena()
#ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetArena() != nullptr
#endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields()
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const HelloResponse& default_instance() {
    return *internal_default_instance();
  }
  static inline const HelloResponse* internal_default_instance() {
    return reinterpret_cast<const HelloResponse*>(
        &_HelloResponse_default_instance_);
  }
  static constexpr int kIndexInFileMessages = 1;
  friend void swap(HelloResponse& a, HelloResponse& b) { a.Swap(&b); }
  inline void Swap(HelloResponse* other) {
    if (other == this) return;
#ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() != nullptr && GetArena() == other->GetArena()) {
#else   // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() == other->GetArena()) {
#endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(HelloResponse* other) {
    if (other == this) return;
    ABSL_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  HelloResponse* New(::google::protobuf::Arena* arena = nullptr) const final {
    return ::google::protobuf::Message::DefaultConstruct<HelloResponse>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const HelloResponse& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom(const HelloResponse& from) { HelloResponse::MergeImpl(*this, from); }

  private:
  static void MergeImpl(
      ::google::protobuf::MessageLite& to_msg,
      const ::google::protobuf::MessageLite& from_msg);

  public:
  bool IsInitialized() const {
    return true;
  }
  ABSL_ATTRIBUTE_REINITIALIZES void Clear() final;
  ::size_t ByteSizeLong() const final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target,
      ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void InternalSwap(HelloResponse* other);
 private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() { return "javaDemo.v1.HelloResponse"; }

 protected:
  explicit HelloResponse(::google::protobuf::Arena* arena);
  HelloResponse(::google::protobuf::Arena* arena, const HelloResponse& from);
  HelloResponse(::google::protobuf::Arena* arena, HelloResponse&& from) noexcept
      : HelloResponse(arena) {
    *this = ::std::move(from);
  }
  const ::google::protobuf::Message::ClassData* GetClassData() const final;

 public:
  ::google::protobuf::Metadata GetMetadata() const;
  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------
  enum : int {
    kResponseFieldNumber = 1,
  };
  // string response = 1;
  void clear_response() ;
  const std::string& response() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_response(Arg_&& arg, Args_... args);
  std::string* mutable_response();
  PROTOBUF_NODISCARD std::string* release_response();
  void set_allocated_response(std::string* value);

  private:
  const std::string& _internal_response() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_response(
      const std::string& value);
  std::string* _internal_mutable_response();

  public:
  // @@protoc_insertion_point(class_scope:javaDemo.v1.HelloResponse)
 private:
  class _Internal;
  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<
      0, 1, 0,
      42, 2>
      _table_;

  static constexpr const void* _raw_default_instance_ =
      &_HelloResponse_default_instance_;

  friend class ::google::protobuf::MessageLite;
  friend class ::google::protobuf::Arena;
  template <typename T>
  friend class ::google::protobuf::Arena::InternalHelper;
  using InternalArenaConstructable_ = void;
  using DestructorSkippable_ = void;
  struct Impl_ {
    inline explicit constexpr Impl_(
        ::google::protobuf::internal::ConstantInitialized) noexcept;
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena);
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena, const Impl_& from,
                          const HelloResponse& from_msg);
    ::google::protobuf::internal::ArenaStringPtr response_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_javaDemo_2fv1_2fjavaDemo_2eproto;
};
// -------------------------------------------------------------------

class HelloRequest final : public ::google::protobuf::Message
/* @@protoc_insertion_point(class_definition:javaDemo.v1.HelloRequest) */ {
 public:
  inline HelloRequest() : HelloRequest(nullptr) {}
  ~HelloRequest() override;
  template <typename = void>
  explicit PROTOBUF_CONSTEXPR HelloRequest(
      ::google::protobuf::internal::ConstantInitialized);

  inline HelloRequest(const HelloRequest& from) : HelloRequest(nullptr, from) {}
  inline HelloRequest(HelloRequest&& from) noexcept
      : HelloRequest(nullptr, std::move(from)) {}
  inline HelloRequest& operator=(const HelloRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline HelloRequest& operator=(HelloRequest&& from) noexcept {
    if (this == &from) return *this;
    if (GetArena() == from.GetArena()
#ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetArena() != nullptr
#endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::google::protobuf::UnknownFieldSet& unknown_fields() const
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.unknown_fields<::google::protobuf::UnknownFieldSet>(::google::protobuf::UnknownFieldSet::default_instance);
  }
  inline ::google::protobuf::UnknownFieldSet* mutable_unknown_fields()
      ABSL_ATTRIBUTE_LIFETIME_BOUND {
    return _internal_metadata_.mutable_unknown_fields<::google::protobuf::UnknownFieldSet>();
  }

  static const ::google::protobuf::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::google::protobuf::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::google::protobuf::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const HelloRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const HelloRequest* internal_default_instance() {
    return reinterpret_cast<const HelloRequest*>(
        &_HelloRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages = 0;
  friend void swap(HelloRequest& a, HelloRequest& b) { a.Swap(&b); }
  inline void Swap(HelloRequest* other) {
    if (other == this) return;
#ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() != nullptr && GetArena() == other->GetArena()) {
#else   // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetArena() == other->GetArena()) {
#endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::google::protobuf::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(HelloRequest* other) {
    if (other == this) return;
    ABSL_DCHECK(GetArena() == other->GetArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  HelloRequest* New(::google::protobuf::Arena* arena = nullptr) const final {
    return ::google::protobuf::Message::DefaultConstruct<HelloRequest>(arena);
  }
  using ::google::protobuf::Message::CopyFrom;
  void CopyFrom(const HelloRequest& from);
  using ::google::protobuf::Message::MergeFrom;
  void MergeFrom(const HelloRequest& from) { HelloRequest::MergeImpl(*this, from); }

  private:
  static void MergeImpl(
      ::google::protobuf::MessageLite& to_msg,
      const ::google::protobuf::MessageLite& from_msg);

  public:
  bool IsInitialized() const {
    return true;
  }
  ABSL_ATTRIBUTE_REINITIALIZES void Clear() final;
  ::size_t ByteSizeLong() const final;
  ::uint8_t* _InternalSerialize(
      ::uint8_t* target,
      ::google::protobuf::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const { return _impl_._cached_size_.Get(); }

  private:
  void SharedCtor(::google::protobuf::Arena* arena);
  void SharedDtor();
  void InternalSwap(HelloRequest* other);
 private:
  friend class ::google::protobuf::internal::AnyMetadata;
  static ::absl::string_view FullMessageName() { return "javaDemo.v1.HelloRequest"; }

 protected:
  explicit HelloRequest(::google::protobuf::Arena* arena);
  HelloRequest(::google::protobuf::Arena* arena, const HelloRequest& from);
  HelloRequest(::google::protobuf::Arena* arena, HelloRequest&& from) noexcept
      : HelloRequest(arena) {
    *this = ::std::move(from);
  }
  const ::google::protobuf::Message::ClassData* GetClassData() const final;

 public:
  ::google::protobuf::Metadata GetMetadata() const;
  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------
  enum : int {
    kRequestFieldNumber = 1,
  };
  // string request = 1;
  void clear_request() ;
  const std::string& request() const;
  template <typename Arg_ = const std::string&, typename... Args_>
  void set_request(Arg_&& arg, Args_... args);
  std::string* mutable_request();
  PROTOBUF_NODISCARD std::string* release_request();
  void set_allocated_request(std::string* value);

  private:
  const std::string& _internal_request() const;
  inline PROTOBUF_ALWAYS_INLINE void _internal_set_request(
      const std::string& value);
  std::string* _internal_mutable_request();

  public:
  // @@protoc_insertion_point(class_scope:javaDemo.v1.HelloRequest)
 private:
  class _Internal;
  friend class ::google::protobuf::internal::TcParser;
  static const ::google::protobuf::internal::TcParseTable<
      0, 1, 0,
      40, 2>
      _table_;

  static constexpr const void* _raw_default_instance_ =
      &_HelloRequest_default_instance_;

  friend class ::google::protobuf::MessageLite;
  friend class ::google::protobuf::Arena;
  template <typename T>
  friend class ::google::protobuf::Arena::InternalHelper;
  using InternalArenaConstructable_ = void;
  using DestructorSkippable_ = void;
  struct Impl_ {
    inline explicit constexpr Impl_(
        ::google::protobuf::internal::ConstantInitialized) noexcept;
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena);
    inline explicit Impl_(::google::protobuf::internal::InternalVisibility visibility,
                          ::google::protobuf::Arena* arena, const Impl_& from,
                          const HelloRequest& from_msg);
    ::google::protobuf::internal::ArenaStringPtr request_;
    mutable ::google::protobuf::internal::CachedSize _cached_size_;
    PROTOBUF_TSAN_DECLARE_MEMBER
  };
  union { Impl_ _impl_; };
  friend struct ::TableStruct_javaDemo_2fv1_2fjavaDemo_2eproto;
};

// ===================================================================




// ===================================================================


#ifdef __GNUC__
#pragma GCC diagnostic push
#pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// -------------------------------------------------------------------

// HelloRequest

// string request = 1;
inline void HelloRequest::clear_request() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.request_.ClearToEmpty();
}
inline const std::string& HelloRequest::request() const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:javaDemo.v1.HelloRequest.request)
  return _internal_request();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void HelloRequest::set_request(Arg_&& arg,
                                                     Args_... args) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.request_.Set(static_cast<Arg_&&>(arg), args..., GetArena());
  // @@protoc_insertion_point(field_set:javaDemo.v1.HelloRequest.request)
}
inline std::string* HelloRequest::mutable_request() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  std::string* _s = _internal_mutable_request();
  // @@protoc_insertion_point(field_mutable:javaDemo.v1.HelloRequest.request)
  return _s;
}
inline const std::string& HelloRequest::_internal_request() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.request_.Get();
}
inline void HelloRequest::_internal_set_request(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.request_.Set(value, GetArena());
}
inline std::string* HelloRequest::_internal_mutable_request() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  return _impl_.request_.Mutable( GetArena());
}
inline std::string* HelloRequest::release_request() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  // @@protoc_insertion_point(field_release:javaDemo.v1.HelloRequest.request)
  return _impl_.request_.Release();
}
inline void HelloRequest::set_allocated_request(std::string* value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.request_.SetAllocated(value, GetArena());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.request_.IsDefault()) {
          _impl_.request_.Set("", GetArena());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:javaDemo.v1.HelloRequest.request)
}

// -------------------------------------------------------------------

// HelloResponse

// string response = 1;
inline void HelloResponse::clear_response() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.response_.ClearToEmpty();
}
inline const std::string& HelloResponse::response() const
    ABSL_ATTRIBUTE_LIFETIME_BOUND {
  // @@protoc_insertion_point(field_get:javaDemo.v1.HelloResponse.response)
  return _internal_response();
}
template <typename Arg_, typename... Args_>
inline PROTOBUF_ALWAYS_INLINE void HelloResponse::set_response(Arg_&& arg,
                                                     Args_... args) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.response_.Set(static_cast<Arg_&&>(arg), args..., GetArena());
  // @@protoc_insertion_point(field_set:javaDemo.v1.HelloResponse.response)
}
inline std::string* HelloResponse::mutable_response() ABSL_ATTRIBUTE_LIFETIME_BOUND {
  std::string* _s = _internal_mutable_response();
  // @@protoc_insertion_point(field_mutable:javaDemo.v1.HelloResponse.response)
  return _s;
}
inline const std::string& HelloResponse::_internal_response() const {
  ::google::protobuf::internal::TSanRead(&_impl_);
  return _impl_.response_.Get();
}
inline void HelloResponse::_internal_set_response(const std::string& value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.response_.Set(value, GetArena());
}
inline std::string* HelloResponse::_internal_mutable_response() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  return _impl_.response_.Mutable( GetArena());
}
inline std::string* HelloResponse::release_response() {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  // @@protoc_insertion_point(field_release:javaDemo.v1.HelloResponse.response)
  return _impl_.response_.Release();
}
inline void HelloResponse::set_allocated_response(std::string* value) {
  ::google::protobuf::internal::TSanWrite(&_impl_);
  _impl_.response_.SetAllocated(value, GetArena());
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
        if (_impl_.response_.IsDefault()) {
          _impl_.response_.Set("", GetArena());
        }
  #endif  // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  // @@protoc_insertion_point(field_set_allocated:javaDemo.v1.HelloResponse.response)
}

#ifdef __GNUC__
#pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)
}  // namespace v1
}  // namespace javaDemo


// @@protoc_insertion_point(global_scope)

#include "google/protobuf/port_undef.inc"

#endif  // GOOGLE_PROTOBUF_INCLUDED_javaDemo_2fv1_2fjavaDemo_2eproto_2epb_2eh